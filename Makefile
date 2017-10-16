.PHONY: proto build

VERSION=$(shell git describe --all --exact-match `git rev-parse HEAD` | grep tags | sed 's/tags\///')
IMAGE_NAME=denkhaus/greeter-micro-srv:$(VERSION)

all: deploy

proto:
	cd proto/hello && protoc --go_out=plugins=micro:. hello.proto

deploy: push
	@IMAGE_NAME=$(IMAGE_NAME) \
	rancher-compose -p services up -d --force-upgrade

push: build
	docker push $(IMAGE_NAME)

build: proto commit
	docker build  -t $(IMAGE_NAME)  .

commit:		
	git add -A && git commit -a -m "autocommit"
	git tag "0.0.$(VERSION)"
	git push origin master

 