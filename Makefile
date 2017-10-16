.PHONY: proto build

VERSION=0.0.13
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
	docker build --build-arg VERSION=$(VERSION) --build-arg GIT_COMMIT=$(shell git rev-list -1 HEAD) -t $(IMAGE_NAME)  .

commit:		
	git add -A && git commit -a -m "autocommit"	
	git push origin master

 