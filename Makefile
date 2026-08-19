BINARY=ccat
SRC=./main.go
VERSION=2.0

build:
	go build -trimpath -x -o $(BINARY) $(SRC)

clean:
	rm -f $(BINARY)

install:
	cp $(BINARY) $(PREFIX)/bin/

test:
	go test ./...

lint:
	golangci-lint run

.PHONY: build clean install test lint
