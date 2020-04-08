IMAGE_NAME ?= $(shell git remote get-url origin | cut -d':' -f2 | sed 's/.git//')
VERSION ?= $(shell git rev-parse --short HEAD)
DOCKER_HOST ?= quay.io

DOCKER_TAG := $(DOCKER_HOST)/$(IMAGE_NAME):$(VERSION)

build:
	docker build -t $(DOCKER_TAG) .

push:
	docker push $(DOCKER_TAG)

.PHONY: build push
