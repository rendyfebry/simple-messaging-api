.DEFAULT_GOAL	:= build

# Commands
SHELL 		:= /bin/bash
GOHOSTOS 	:= $(shell go env GOHOSTOS)
RM				:= rm -rf
MKDIR			:= mkdir -p
TAR				:= tar
GOTEST		:= go test -v
SETENV		:= export
LINT			:= golint

# Options
CGO_ENABLED	:= 0
DATADIR			:= data
DOCKERDIR		:= docker
RELEASEDIR	:= release
BINDIR			:= bin
TESTDIR     := tests
BINARY			:= msgapi

ifndef GOOS
    GOOS := $(GOHOSTOS)
endif

ifndef GOARCH
	GOARCH := $(shell go env GOHOSTARCH)
endif

GOFILES		= $(shell find . -type f -name '*.go' -not -path "./vendor/*")
GODIRS		= $(shell go list -f '{{.Dir}}' ./...)
GOPKGS		= $(shell go list ./... | grep -v /vendor/)

install:
	dep ensure -v

test:
	$(GOTEST) ./...

build-quick:
	@${MKDIR} ${BINDIR}
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o './$(BINDIR)/$(BINARY)' 

build: install build-quick
