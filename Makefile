

include ./hack/help.mk


UID:=$(shell id -u)
GID:=$(shell id -g)
PWD:=$(shell pwd)


.PHONY: setup
setup: ##@setup builds the container image(s) and starts the setup
	docker-compose --compatibility build
	docker-compose --compatibility up -d


.PHONY: stop
stop: ##setup stops the setup (e.g. running containers)
	docker-compose --compatibility stop -t 2


.PHONY: clean
clean: ##@setup clean setup
	docker-compose --compatibility down -t 2


.PHONY: status
status: ##@telemetry shows the state of the running setup
	docker-compose --compatibility ps


.PHONY: logs
logs: ##@telemetry shows logs
	docker-compose --compatibility logs


.PHONY: cli
cli: ##@dev exec's into container with all the dev tools
	# docker run -it --rm -u $(UID):$(GID) -v $(PWD):/source -w /source golang:1.10.3 bash
	# docker run -it --rm -v $(PWD):/go/src/go-start -w /go/src/go-start -v $(PWD)/certs:/certs -p 8083:8083 golang:1.13.5 bash
	docker-compose --compatibility exec goenv bash


.PHONY: mkdocker
mkdocker: ##@minikube reuse minikube docker env to run docker cmd's | e.g. `make mkdocker ARGS="ps"`
	@eval $$(minikube docker-env) ;\
	docker $(ARGS)


.PHONY: env
env: ##@setup cp .env.template to .env
	cp .env.template .env

