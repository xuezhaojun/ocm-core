SHELL :=/bin/bash

all: build
.PHONY: all

# Include the library makefile
include $(addprefix ./vendor/github.com/openshift/build-machinery-go/make/, \
	golang.mk \
	targets/openshift/deps.mk \
	targets/openshift/images.mk \
	targets/openshift/kustomize.mk \
	lib/tmp.mk \
)

IMAGE_REGISTRY?=quay.io/open-cluster-management
IMAGE_TAG?=latest
IMAGE_NAME?=$(IMAGE_REGISTRY)/ocm:$(IMAGE_TAG)

# This will call a macro called "build-image" which will generate image specific targets based on the parameters:
# $0 - macro name
# $1 - target suffix
# $2 - Dockerfile path
# $3 - context directory for image build
# It will generate target "image-$(1)" for building the image and binding it as a prerequisite to target "images".
$(call build-image,ocm,$(IMAGE_REGISTRY)/ocm:$(IMAGE_TAG),./Dockerfile,.)

build-e2e:
	go test -c -o e2e.test ./test/e2e

test-e2e: build-e2e
	./e2e.test