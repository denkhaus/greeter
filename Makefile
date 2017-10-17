.PHONY: proto build

VERSION=$(shell git semver get)
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
	git add -A
	-@if [ $(shell git status --porcelain 2>/dev/null | egrep "^(M| M)" | wc -l) ]; then \
		git semver next 2>/dev/null && git commit -a -m "proceed" && git push origin master; \
	fi