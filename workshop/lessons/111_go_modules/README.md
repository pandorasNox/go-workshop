# go modules

go mod init

`go.mod`
`go.sum`

- non resouce module names
    - `mymodule`
    - go get doesn't work anymore
- "normal" resource modules names
    - `github.com/pandorasnox/go-workshop/workshop/lessons/111_go_modules`
    - can be installed via go get (when resource exists under uri)

- `go mod init`

- `go mod vendor` make vendored copy of dependencies

- path
    - 


- module name
  - or as I would call it, the source from where you would fetch the dependencie
  - alternativly you can choose that this is a package which never will be a fetchable resource
    - you can give it any name, e.g. `mymodlule`
    - but with this go get doesn't work anymore
    - however this will shortn the import path
      - e.g. `import "mymodule/apackage"`

----

## all you need to know about go modules (and a little more) in order to work with it

- dependency management
- you want to be able to import some or all packages from one of your modules into another module
- there is a special "private" kind of modules, these can't be included from other modules
  - to declare a module as "private", you give it a module name which isn't reachable as a resource
- with go modules you can put the directory of the module locally whereever you want
- most common resource standard is providing the resource as a git repository
  - other ways are possible
- a module is a directory in which
  - contains a go.mod file
  - is as an upstream resource reachable (e.g. github.com/someuser/somemodule)
  - can contain a go.sum file
    - is meant for the dependencies of this module
- a module can contain
  - one or multimple packages (which lay in subdirectories or, as it is one, directly in the root)
- you can have multiple modules in a single repository
  - again, the important point is the modules are reachable as a resource
  - although, it's not recommended
  > For all but power users, you probably want to adopt the usual convention that one repo = one module. It's important for long-term evolution of code storage options that a repo can contain multiple modules, but it's almost certainly not something you want to do by default.
- what a reachable resource is:
  - 
- packages are imported using the full path including the module path
  - you don't import modules, you always import packages of these modules
- in the GOPATH go modules are per default disabled
  - >if you want to work in your GOPATH
  ```
  $ export GO111MODULE=on                         # manually active module mode
  $ cd $GOPATH/src/<project path>                 # e.g., cd $GOPATH/src/you/hello
  ```
- 'go.sum' is _not_ a lock file
  - https://github.com/golang/go/wiki/Modules#is-gosum-a-lock-file-why-does-gosum-include-information-for-module-versions-i-am-no-longer-using
  - go.mod files in a build provide enough information for 100% reproducible builds
  - >No matter the source of the modules, the go command checks downloads against known checksums, to detect unexpected changes in the content of any specific module version from one day to the next. This check first consults the current module's go.sum file but falls back to the Go checksum database, controlled by the GOSUMDB and GONOSUMDB environment variables.
- go modules uses semver, it's required
  - also, versioning is hard
    - v2 versions of modules are treated special, v2 versions affact the import path, the v2 has to be part of the path than

example:
module A
  import "B"
  import "C"
  indirect import D (v2)
module B
  imports D (v2)
module C
  imports D (v3)
module D
  version 1
  version 2
  version 3


### more go module infos
- https://github.com/golang/go/wiki/Modules
- https://stackoverflow.com/questions/52125437/go-module-init-without-vcs-git-fails-with-cannot-determine-module-path
- https://medium.com/rungo/anatomy-of-modules-in-go-c8274d215c16
- https://golang.org/cmd/go/#hdr-Module_proxy_protocol
