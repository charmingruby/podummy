DOCKER_IMAGE_NAME = podummy

###################
# Docker          #
###################
.PHONY: docker-build
docker-build:
	docker build -t ${DOCKER_IMAGE_NAME} .

.PHONY: docker-run
docker-run:
	docker run -p 3000:3000 ${DOCKER_IMAGE_NAME}
