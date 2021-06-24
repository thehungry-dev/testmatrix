.PHONY: clean build deps

build:
	go build

clean:
	go clean -i github.com/thehungry-dev/testmatrix...

deps:
	go get && go mod tidy
