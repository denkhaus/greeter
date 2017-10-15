.PHONY: proto build

all: deploy

proto:
	for d in proto; do \
		for f in $$d/**/proto/*.proto; do \
			protoc --go_out=plugins=micro:. $$f; \
			echo compiled: $$f; \
		done \
	done

deploy: push
	@rancher-compose -p services up -d --force-upgrade

push: build
	docker push denkhaus/greeter:latest

build: proto commit
	docker build  -t denkhaus/greeter  .

commit:
	git add -A && git commit -a -m "proceed"

