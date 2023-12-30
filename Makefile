IMAGE_NAME = joerx/hellod
VERSION ?= $(shell git rev-parse --short HEAD)
DOCKER_HOST ?= ghcr.io

DOCKER_REPO := $(DOCKER_HOST)/$(IMAGE_NAME)

default: clean build

build:
	go build -o out/hellod .

clean:
	rm -rf out

docker-build:
	docker build -t $(DOCKER_REPO):$(VERSION)-arm64 --build-arg ARCH=arm64 .

docker-push: docker-build
	docker push $(DOCKER_REPO):$(VERSION)-arm64

.PHONY: build push docker-build docker-push default
