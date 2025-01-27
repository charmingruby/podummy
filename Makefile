DOCKER_IMAGE_NAME = podummy
VERSION ?= blue
SERVER_PORT ?= 3000

###################
# Docker          #
###################
.PHONY: docker-build
docker-build:
	docker build --build-arg VERSION=$(VERSION) --build-arg SERVER_PORT=$(SERVER_PORT) -t $(DOCKER_IMAGE_NAME):$(VERSION) .

.PHONY: docker-run
docker-run:
	docker run -p $(SERVER_PORT):$(SERVER_PORT) $(DOCKER_IMAGE_NAME):$(VERSION)
