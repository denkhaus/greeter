all: deploy



deploy: build


build:
	docker build  -t denkhaus/greeter .

