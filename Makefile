.PHONY: default build test lint help

# Default target to build the project
default: build

help:	## List all available commands
	@echo 'Available commands:'
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

test:	## Run all tests
	go test ./... -count=1

build:	## Build the project
	go build ./...