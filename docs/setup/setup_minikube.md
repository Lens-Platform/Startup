## Setting up minikube for local development and testing of blackspace services
minikube is local Kubernetes, focusing on making it easy to learn and develop for Kubernetes.

All you need is Docker (or similarly compatible) container or a Virtual Machine environment, and Kubernetes is a single command away: minikube start

### What you need
What you’ll need
* 2 CPUs or more
* 2GB of free memory
* 20GB of free disk space
* Internet connection
* Container or virtual machine manager, such as: Docker, Hyperkit, Hyper-V, KVM, Parallels, Podman, VirtualBox, or VMWare

### Installation
#### linux

#### macos
If the Brew Package Manager installed:

```bash
brew install minikube
which minikube
```
If `which minikube` fails after installation via brew, you may have to remove the minikube cask and link the binary:

```bash
brew cask remove minikube
brew link minikube
```

Otherwise, download minikube directly:

```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64
sudo install minikube-darwin-amd64 /usr/local/bin/minikube
```

#### Windows
##### Windows Package Manager
If the [Windows Package Manager](https://docs.microsoft.com/en-us/windows/package-manager/) is installed, use the following command to install minikube:

```bash
winget install minikube
```

##### Chocolatey
If the [Chocolatey Package Manager](Chocolatey Package Manager) is installed, use the following command:
```bash
choco install minikube
```

Stand-alone Windows Installer
Otherwise, download and run the [Windows installer](https://storage.googleapis.com/minikube/releases/latest/minikube-installer.exe)

#### Linux
For Linux users, they are 3 easy download options (for each architecture):

##### x86
###### Binary
```bash
 curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
 sudo install minikube-linux-amd64 /usr/local/bin/minikube
```
##### Debian
```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_latest_amd64.deb
sudo dpkg -i minikube_latest_amd64.deb
```
##### RPM
```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-latest.x86_64.rpm
sudo rpm -ivh minikube-latest.x86_64.rpm
```

##### ARM
###### Binary
```bash
 curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-arm64
 sudo install minikube-linux-arm64 /usr/local/bin/minikube
```
##### Debian
```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_latest_arm64.deb
sudo dpkg -i minikube_latest_arm64.deb
```
##### RPM
```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-latest.aarch64.rpm
sudo rpm -ivh minikube-latest.aarch64.rpm
```

### Start cluster
From a terminal with administrator access (but not logged in as root), run:

```bash
minikube start
```

to start with a driver (hyperkit, hyperv, kvm2, docker, none, podman, parallels, virtualbox, vmware)

```bash
minikube start --driver=providername
```

to set the provider as a the default
```bash
minikube config set driver providername
```

If minikube fails to start, see the [drivers page](https://minikube.sigs.k8s.io/docs/drivers/) for help setting up a compatible container or virtual-machine manager.

### Interact with your cluster
If you already have kubectl installed, you can now use it to access your shiny new cluster:
```bash
kubectl get po -A
```

Alternatively, minikube can download the appropriate version of kubectl, if you don’t mind the double-dashes in the command-line:
```bash
minikube kubectl -- get po -A
```

Initially, some services such as the storage-provisioner, may not yet be in a Running state. This is a normal condition during cluster bring-up, and will resolve itself momentarily. For additional insight into your cluster state, minikube bundles the Kubernetes Dashboard, allowing you to get easily acclimated to your new environment:

```bash
minikube dashboard
```

### Deploy applications
Create a sample deployment and expose it on port 8080:
```bash
kubectl create deployment hello-minikube --image=k8s.gcr.io/echoserver:1.4
kubectl expose deployment hello-minikube --type=NodePort --port=8080
```
It may take a moment, but your deployment will soon show up when you run:

```bash
kubectl get services hello-minikube
```

The easiest way to access this service is to let minikube launch a web browser for you:

```bash
minikube service hello-minikube
```

Alternatively, use kubectl to forward the port:
```bash
kubectl port-forward service/hello-minikube 7080:8080
```

Tada! Your application is now available at http://localhost:7080/

To access a LoadBalancer deployment, use the “minikube tunnel” command. Here is an example deployment:

```bash
kubectl create deployment balanced --image=k8s.gcr.io/echoserver:1.4
kubectl expose deployment balanced --type=LoadBalancer --port=8080
```

In another window, start the tunnel to create a routable IP for the ‘balanced’ deployment:

```bash
minikube tunnel
```

To find the routable IP, run this command and examine the EXTERNAL-IP column:

```bash
kubectl get services balanced
```

Your deployment is now available at <EXTERNAL-IP>:8080

### Managing the cluster
Pause Kubernetes without impacting deployed applications:

```bash
minikube pause
```

Halt the cluster
```bash
minikube stop
```

Increase the default memory limit (requires a restart)
```bash
minikube config set memory 16384
```

Browse the catalog of easily install kuberneted service
```bash
minikube addons list
```

Create a second cluster running an older kubernetes running an older release
```bash
minikube start -p aged --kubernetes-version=v1.16.1
```

Delete all minikube clusters
```bash
minikube delete --all
```

### Tutorials
[references](https://minikube.sigs.k8s.io/docs/tutorials/)
[command reference](https://minikube.sigs.k8s.io/docs/commands/)
[minikube handbook](https://minikube.sigs.k8s.io/docs/handbook/)
