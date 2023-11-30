MAKEFLAGS += --silent

all: clean format test build

## clean: Removes any previously created build artifacts.
clean:
	rm -rf ./bin/

xk6:
	go install go.k6.io/xk6/cmd/xk6@latest

## build: Builds a custom 'k6' with the local extension.
build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux xk6 build --with $(shell go list -m)=.
	upx -9 ./k6
	mkdir ./bin/
	cp k6 ./bin/
	rm -f ./k6

## format: Applies Go formatting to code.
format:
	go fmt ./...

## test: Executes any unit tests.
test:
	go test -cover -race ./...

.PHONY: build clean format help test
