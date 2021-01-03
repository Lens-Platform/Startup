## Skaffold
Skaffold is useful as it simplifies your development workflow by organizing common development stages into one simple command. Everytime `skaffold
 dev` is ran the system does the following

 * Collects and watches your source code for changes
 * Syncs files directly to pods if user marks them as syncable
 * Builds artifacts from the source code
 * Tests the built artifacts using container-structure-tests
 * Tags the artifacts
 * Pushes the artifacts
 * Deploys the artifacts
 * Monitors the deployed artifacts
 * Cleans up deployed artifacts on exit (Ctrl+C)


### Installation
#### Linux
The latest stable binary can be found here:

https://storage.googleapis.com/skaffold/releases/latest/skaffold-linux-amd64

Simply download it and add it to your PATH. Or, copy+paste this command in your terminal:

```bash
curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-linux-amd64 && \
sudo install skaffold /usr/local/bin/
```

#### Macos
##### Binary
The latest stable binary can be found here:

https://storage.googleapis.com/skaffold/releases/latest/skaffold-darwin-amd64

Simply download it and add it to your PATH. Or, copy+paste this command in your terminal:

```bash
curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-darwin-amd64 && \
sudo install skaffold /usr/local/bin/
```
##### Homebrew
```bash
brew install skaffold
```

##### MacPorts
```bash
sudo port install skaffold
```


To setup skaffold, please follow the steps outlined at the [documentation website](https://skaffold.dev/docs/)

#### Windows
The latest stable release binary can be found here:

     https://storage.googleapis.com/skaffold/releases/latest/skaffold-windows-amd64.exe

     Simply download it and place it in your PATH as skaffold.exe.
##### Chocolatey
```bash
choco install -y skaffold
```

#### Docker
Stable binary
For the latest stable release, you can use:

```bash
docker run gcr.io/k8s-skaffold/skaffold:latest skaffold <command>
```

