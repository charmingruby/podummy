DOCKER_IMAGE_NAME = podummy
VERSION ?= blue
SERVER_PORT?=3000

###################
# Docker          #
###################
.PHONY: docker-build
docker-build:
	docker build --build-arg VERSION=$(VERSION) -t $(DOCKER_IMAGE_NAME):$(VERSION) .

.PHONY: docker-run
docker-run:
	docker run -p $(SERVER_PORT):8080 $(DOCKER_IMAGE_NAME):$(VERSION)
