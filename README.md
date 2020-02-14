
## vscode remote container setup
- install the vscode remote container extension
- run `make setup` to start docker-compose setup
- click on the bottom bar on the left "Open a remote window" button (looks almost like "><")
- choose to attach to running container ("Remote-Containers: Attach to Running Container...") and choose the one with the service name in the conatiner name e.g. when the docker-compose service name is `goenv` the container name should be something like `go-start_goenv_1`
- if needed open via vsc cmd palette `Remote-Containers: Open Attached Container Configuration File` [more info](https://code.visualstudio.com/docs/remote/containers#_attached-container-configuration-files)
