FROM golang:1.13.5 AS dev

# Avoid warnings by switching to noninteractive
ENV DEBIAN_FRONTEND=noninteractive

# This Dockerfile adds a non-root user with sudo access. Use the "remoteUser"
# property in devcontainer.json to use it. On Linux, the container user's GID/UIDs
# will be updated to match your local UID/GID (when using the dockerFile property).
# See https://aka.ms/vscode-remote/containers/non-root-user for details.
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=${USER_UID}

# add new group `golang`, add root
# permit golang group for /go dirctory
RUN groupadd --gid 3333 golang
RUN usermod -a -G golang root

# Configure apt, install packages and tools
RUN apt-get update \
    && apt-get -y install --no-install-recommends apt-utils dialog 2>&1 \
    #
    # Verify git, process tools, lsb-release (common in install instructions for CLIs) installed
    && apt-get -y install git iproute2 procps lsb-release \
    # # Install gocode-gomod
    # && go get -x -d github.com/stamblerre/gocode 2>&1 \
    # && go build -o gocode-gomod github.com/stamblerre/gocode \
    # && mv gocode-gomod $GOPATH/bin/ \
    #
    # Install Go tools
    && go get -u -v \
        # github.com/mdempsky/gocode \
        # like `go list all`
        github.com/uudashr/gopkgs/cmd/gopkgs \
        github.com/ramya-rao-a/go-outline \
        github.com/acroca/go-symbols \
        github.com/godoctor/godoctor \
        golang.org/x/tools/cmd/guru \
        golang.org/x/tools/cmd/gorename \
        github.com/rogpeppe/godef \
        github.com/zmb3/gogetdoc \
        # github.com/haya14busa/goplay/cmd/goplay \
        github.com/sqs/goreturns \
        github.com/josharian/impl \
        github.com/davidrjenni/reftools/cmd/fillstruct \
        github.com/fatih/gomodifytags \
        github.com/cweill/gotests/... \
        golang.org/x/tools/cmd/goimports \
        golang.org/x/lint/golint \
        # golang.org/x/tools/cmd/gopls \
        github.com/alecthomas/gometalinter \
        honnef.co/go/tools/... \
        github.com/golangci/golangci-lint/cmd/golangci-lint \
        # golint replacement
        github.com/mgechev/revive \
        # debugger
        github.com/derekparker/delve/cmd/dlv \
        2>&1

ENV GO111MODULE on
RUN go get golang.org/x/tools/gopls@v0.3.2

    # Create a non-root user to use if preferred - see https://aka.ms/vscode-remote/containers/non-root-user.
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd -s /bin/bash --uid $USER_UID --gid $USER_GID -m $USERNAME \
    # [Optional] Add sudo support
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

#  add vcode user to group `golang`
RUN usermod -a -G golang ${USERNAME}
RUN chown -R :golang /go
RUN chmod -R g+rwx /go
RUN chown -R :golang /etc/ssh
RUN chmod -R g+rwx /etc/ssh

RUN apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/* /go/src

# Switch back to dialog for any ad-hoc use of apt-get
ENV DEBIAN_FRONTEND=dialog

USER vscode
RUN mkdir -p /home/vscode/.vscode-server/extensions
USER root

COPY build/build-image/rootfs /

# USER vscode
