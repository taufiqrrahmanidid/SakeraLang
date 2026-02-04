.PHONY: build run clean install test

build:
	go mod tidy
	go build -o sakera ./cmd/sakera

run:
	./sakera

clean:
	rm -f sakera

install:
	go install ./cmd/sakera

test:
	./sakera examples/hello.sakera

all: clean build
