# vulcan_ramp_up
Cloud Development Refresher

### Start Server
```bash
$ go run hello.go
```

### GET request
```bash
$ curl http://localhost/hello
```

### Run Unit Testing
```bash
$ go test
```

## What is a container?
A container is a standard unit of software that packages up code and all its dependencies so the application runs quickly and reliably from one computing environment to another

## What are the advantages of containers?
A VM image includes of a complete operating system. A container image in contrast assumes most of the operating system capabilities from the host operating system on the machine.

Therefore, a container image is much more lightweight in comparison to a VM image. It takes up less space on disk/memory, and it can be started and stopped much more rapidly, because another operating system is not initiated and started.

## What is Docker?
Docker is an open platform for developing, shipping, and running applications. Docker enables you to separate your applications from your infrastructure so you can deliver software quickly.

Docker provides the ability to package and run an application in a loosely isolated environment called a container. The isolation and security allows you to run many containers simultaneously on a given host. 

### Build a Container
```bash
docker build -t go-gin-docker ./vulcan_ramp_up
```

### Run a Container
```bash
docker run -p 80:80 go-gin-docker
```

## What is Kubernetes?
Kubernetes is an open source container orchestration platform that automates many of the manual processes involved in deploying, managing, and scaling containerized applications.

Users can cluster together groups of hosts running containers, and Kubernetes helps you easily and efficiently manage those clusters.

### Create a Kubernetes Cluster
```bash
k3d cluster create --config ./k3d/cluster_create.yaml
```

### Deploy hello-world to the cluster
```bash
kubectl apply --filename k8s/
```

## What is Helm?
Helm is a Kubernetes deployment tool for automating creation, packaging, configuration, and deployment of applications and services to Kubernetes clusters.

### Install using Helm Chart
```bash
helm install hello-world-chart .
```


