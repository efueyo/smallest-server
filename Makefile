.PHONY: build
build:
	docker build -t smallest-server .

.PHONY: run
run:
	go run .

