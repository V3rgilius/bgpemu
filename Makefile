GOPATH ?= ${HOME}/go
BGPEMU_CLI_BIN := bgpemu
INSTALL_DIR := /usr/local/bin

## Run unit tests
test:
	go test ./...

## Targets below are for integration testing only

.PHONY: build
## Build kne
build:
	CGO_ENABLED=0 go build -o $(BGPEMU_CLI_BIN) -ldflags="-s -w" ./main.go

.PHONY: install
## Install kne cli binary to user's local bin dir
install: build
	sudo mv $(BGPEMU_CLI_BIN) $(INSTALL_DIR)
 