IMAGE_NAME ?= $(shell git remote get-url origin | cut -d':' -f2 | sed 's/.git//')
VERSION ?= $(shell git rev-parse --short HEAD)
DOCKER_HOST ?= ghcr.io

DOCKER_TAG := $(DOCKER_HOST)/$(IMAGE_NAME):$(VERSION)

default: clean build

build:
	go build -o out/hellod .

clean:
	rm -rf out

docker-build:
	docker build -t $(DOCKER_TAG) .

docker-push: docker-build
	docker push $(DOCKER_TAG)

.PHONY: build push docker-build docker-push default
