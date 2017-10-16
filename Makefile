.PHONY: proto build

VERSION=0.0.3
IMAGE_NAME=denkhaus/greeter-micro-svc

all: deploy

proto:
	cd proto/hello && protoc --go_out=plugins=micro:. hello.proto

deploy: push
	@IMAGE_NAME=$(IMAGE_NAME) \
	rancher-compose -p services up -d --force-upgrade

push: build
	docker push $(IMAGE_NAME):latest

build: proto commit
	docker build  -t $(IMAGE_NAME)  .

commit:
	if [ $(shell git describe --all --exact-match `git rev-parse HEAD` | grep tags | sed 's/tags\///') != $(VERSION) ]; \
		then git tag $(VERSION); \
	fi
	git add -A && git commit -a -m "autocommit"
	git push origin master

