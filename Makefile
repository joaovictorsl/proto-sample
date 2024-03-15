.PHONY: run build

build:
	go build -o ./bin/proto-sample

run: build
	./bin/proto-sample
