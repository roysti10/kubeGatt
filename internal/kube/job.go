package kube

import (
	"context"
	"fmt"

	"github.com/roysti10/termCI/internal/structs"
	"github.com/roysti10/termCI/internal/util"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	//"strings"
)

func CreateJob(clientset *kubernetes.Clientset, pipeline *structs.Pipeline, stageName string) error {
	var stepNames []string
	stepNames = pipeline.Stages[stageName]
	var containers []corev1.Container
	for _, stepName := range stepNames {
		var step structs.Step
		step = pipeline.Steps[stepName]
		//commands := strings.Split(step.Command, "\n")
		container := corev1.Container{
			Name:       stepName,
			Image:      step.Image,
			WorkingDir: "/mnt/repo",
			Command:    []string{"sh", "-c", step.Command},
			VolumeMounts: []corev1.VolumeMount{
				{
					Name:      "repo",
					MountPath: "/mnt/repo",
				},
			},
		}
		containers = append(containers, container)
	}
	url, err := util.GetGitCloneUrl()
	if err != nil {
		return err
	}
	fmt.Println(url)
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      stageName,
			Namespace: "default",
		},
		Spec: batchv1.JobSpec{
			BackoffLimit: int32Ptr(0),
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					RestartPolicy: corev1.RestartPolicyNever,
					Volumes: []corev1.Volume{
						{
							Name: "repo",
							VolumeSource: corev1.VolumeSource{
								EmptyDir: &corev1.EmptyDirVolumeSource{},
							},
						},
					},
					InitContainers: []corev1.Container{
						{
							Name:    "repo-clone",
							Image:   "alpine/git",
							Command: url,
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "repo",
									MountPath: "/mnt/repo",
								},
							},
						},
					},
					Containers: containers,
				},
			},
		},
	}
	jobClient := clientset.BatchV1().Jobs("default")
	_, err = jobClient.Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Job %s created successfully!\n", stageName)
	return nil
}

func int32Ptr(i int32) *int32 { return &i }
