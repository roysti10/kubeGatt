package run

import (
	"fmt"
	"github.com/roysti10/termCI/internal/kube"
	"github.com/roysti10/termCI/internal/structs"
	"github.com/roysti10/termCI/internal/util"
	"k8s.io/client-go/kubernetes"
	"log"
	"sync"
)

func worker(clientset *kubernetes.Clientset, pipeline *structs.Pipeline, stages <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for stageName := range stages {
		wg.Add(1)
		go func(stage string) {
			defer wg.Done()
			if err := kube.CreateJob(clientset, pipeline, stage); err != nil {
				log.Printf("Error while creating job for stage %s: %v", stage, err)
			}
		}(stageName)
	}
}

func LaunchStagesWithWorkerPool(clientset *kubernetes.Clientset, pipeline *structs.Pipeline, workerCount int) {
	stages := make(chan string, len(pipeline.Stages))
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(i)
		go worker(clientset, pipeline, stages, &wg)
	}
	for stage, _ := range pipeline.Stages {
		stages <- fmt.Sprintf("%s", stage)
	}
	close(stages)
	wg.Wait()
}

func Execute() error {
	fmt.Printf("hello")
	var pipeline *structs.Pipeline
	pipeline, err := util.Validate()
	if err != nil {
		return err
	}
	var clientset *kubernetes.Clientset
	clientset, err = kube.Config()
	if err != nil {
		return err
	}

	LaunchStagesWithWorkerPool(clientset, pipeline, 3)
	return nil
}
