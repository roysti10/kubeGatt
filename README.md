# KubeGatt

A ci/cd fraemwork which runs stages as kubernetes jobs

Note: This is a very raw implementation and I will try to improve it as and when I get time

## Inspiration

This project was inspired by [Fregatt](https://medium.com/pipedrive-engineering/how-did-we-built-our-own-ci-cd-framework-from-scratch-part-ii-d42557aa2641), in their project they used sibling containers and I attempted to replicate the same using jobs

## How to run

Create a yaml file and save it inside the .term directory

*Example yaml file*
```yaml
stages:
  stage1:
    - step1
  stage2:
    - step1

steps:
  step1:
    image: golang 
    command: |
      echo "hello world"
      go install golang.org/x/lint/golint@latest
      golint ./...

```

Each stage runs parallely and each step in a stage runs sequentially

Build the application:

`go build -o term ./cmd`

This will generate a executable which can be run as `./term run`

Prerequisties: a minikube environment running

To see logs you can use: `kubectl logs -l job-name=stage1`


## License

MIT License




