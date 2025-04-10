IMAGE=signpost-example

build-and-run: build run

build:
	docker build -t ${IMAGE} .

run:
	docker run -p 8000:8000 ${IMAGE}
