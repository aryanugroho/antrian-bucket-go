#!/bin/bash

export PKGS=$(shell go list ./...)

all:
	@make build && make gotest

gotest:
	@go test -v --cover --race ${PKGS}

build:
	@echo "building.."
	@go build -o ./antrian-bucket ./cmd/antrian-bucket/
	@echo "done."

docker:
	@echo "building docker..."
	@docker build -t antrian-bucket .
	@echo "done."

