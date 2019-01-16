## Aladdin: a maxflow model based scheduler

Aladdin is a new apporach to cluster scheduling that models the problem of scheduling as a network flow optimization problem. Aladdin aims on providing:

- Max-number tasks scheduling
- Lower latency on scheduling decisions
- Inter tasks constraint support

Aladdin is an Go implementation of the scheduler. The goal of this project is to integrate Aladdin in [Kubernetes](https://github.com/kubernetes/kubernetes) as an alternative scheduler.

## Current State of Project:

The project so far is an early stage prototype with max-number tasks scheduling and Kubernetes integrate implemented. Aladdin uses the Kubernetes API allowing it to batch schedule pods. Currently the implementation has no support for inter tasks constraint, and we will implement it for the foreseeable future.

## Trying it Out:

To get the scheduler up and running there are xxx steps to currently test it out.

### Running on a live kubernetes cluster

#### Step 1: 

You need to clone this repo, and move to `<RepoPath>/cmd/kube-scheduler/`, and build this repo with command`env GOOS=linux GOARCH=amd64 go build scheduler.go`

#### Step 2:

You can test the scheduler by running it inside of a container on the kubernetes master node. You can build the image from `build/Dockerfile` yourself, before that you need to move excutable scheduler binary file to /build path, then build image with command 

```bash
docker build -t aladdin-scheduler:1.0 .
docker --push larryyang/aladdin-scheduler:1.0
```

#### Step 3:

We created a Deployment configuration file and ran it in an existing Kubernetes cluster, using the Deployment resource rather than creating a Pod resource directly because the Deployment resource can better handle the case of a scheduler running node failure. Here is the Deployment configuration example, saved as the `aladdin-scheduler.yaml` file:

```yaml
apiVersion: v1
kind: Deployment
metadata: 
 labels:
  component: scheduler
  tier: control-plane
 name: custom-scheduler
 namespace: kube-system
spec:
  selector:
    matchLabels:
      component: scheduler
      tier: control-plane
  replicas: 1
  template:
    metadata:
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ""
creationTimestamp: null
      labels:
        component: scheduler
        tier: control-plane
      name: aladdin-scheduler
      namespace: kube-system
    spec:
      containers:
        - command:
          - /usr/local/bin/aladdin-scheduler
          - --address=0.0.0.0
          - --scheduler-name=aladdin-scheduler
          - --kubeconfig=/etc/kubernetes/scheduler.conf
          - --leader-elect=false
          - --port=10253
          image: aladdin-scheduler:1.0
          imagePullPolicy: IfNotPresent
          livenessProbe:
            failureThreshold: 8
            httpGet:
              path: /healthz
              port: 10253
              scheme: HTTP
            initialDelaySeconds: 15
            timeoutSeconds: 15
          name: custom-scheduler
          resources:
            requests:
              cpu: 100m
            volumeMounts:
              - mountPath: /etc/kubernetes/scheduler.conf 
                name: kubeconfig
                readOnly: true
      hostNetwork: true
      priorityClassName: system-cluster-critical 
      volumes:
        - hostPath:
            path: /etc/kubernetes/scheduler.conf
            type: FileOrCreate
          name: kubeconf
```

#### Step 4

Create Deployment resource in Kubernetes cluster

`kubectl create -f aladdin-scheduler.yaml`