.PHONY: proto build

IMAGE_NAME=denkhaus/greeter-micro-svc

all: deploy

proto:
	cd proto/hello && protoc --go_out=plugins=micro:. hello.proto

deploy: push
	@IMAGE_NAME=$(IMAGE_NAME) \
	@rancher-compose -p services up -d --force-upgrade

push: build
	docker push $(IMAGE_NAME):latest

build: proto commit
	docker build  -t $(IMAGE_NAME)  .

commit:
	git add -A && git commit -a -m "proceed"
	git push origin master

