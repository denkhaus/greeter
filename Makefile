.PHONY: proto build

all: deploy

proto:
	cd proto/hello && protoc --go_out=plugins=micro:. hello.proto

deploy: push
	@rancher-compose -p services up -d --force-upgrade

push: build
	docker push denkhaus/greeter:latest

build: proto commit
	docker build  -t denkhaus/greeter  .

commit:
	git add -A && git commit -a -m "proceed"
	git push origin master

