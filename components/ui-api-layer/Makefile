IMG_NAME = ui-api-layer
IMG = $(DOCKER_PUSH_REPOSITORY)$(DOCKER_PUSH_DIRECTORY)/$(IMG_NAME)
TAG = $(DOCKER_TAG)

ci-pr: resolve build-and-test build-image push-image
ci-master: resolve build-and-test build-image push-image
ci-release: resolve build-and-test build-image push-image

resolve:
	dep ensure -v -vendor-only
build-and-test:
	./before-commit.sh ci
build-image:
	docker build -t $(IMG_NAME) .
push-image:
	docker tag $(IMG_NAME) $(IMG):$(TAG)
	docker push $(IMG):$(TAG) 