all: deploy


deploy: push
	@rancher-compose -p services up -d --force-upgrade

push: build
	docker push denkhaus/greeter:latest

build: commit
	docker build  -t denkhaus/greeter  .

commit:
	git add -A && git commit -a -m "proceed"

