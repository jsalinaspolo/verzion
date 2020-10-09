#!/usr/bin/env make

SHELL := /bin/bash
.SHELLFLAGS = -ec

.PHONY: run
run:
	go run main.go
