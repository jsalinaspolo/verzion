#!/usr/bin/env make

SHELL := /bin/bash
.SHELLFLAGS = -ec

.PHONY: run
run:
	go run main.go

.PHONY: install
install:
	go install

.PHONY: test
test:
	go test -cover ./...
