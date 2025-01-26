.PHONY: all build start build-client build-server apply

SERVER_IMAGE = server
CLIENT_IMAGE = client
IMAGE_TAG ?= latest

start:
	minikube start

build: build-server build-client

build-server: start
	eval $$(minikube -p minikube docker-env) && docker build -f Dockerfile.server -t $(SERVER_IMAGE):$(IMAGE_TAG) .

build-client: start
	eval $$(minikube -p minikube docker-env) && docker build -f Dockerfile.client -t $(CLIENT_IMAGE):$(IMAGE_TAG) .

clean:
	minikube delete

all: clean build apply

apply:
	kubectl apply -f ./k8s

