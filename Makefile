ifeq ($(OS), Windows_NT)
	VERSION := $(shell git describe --exact-match --tags 2>nil)
	HOME := $(HOMEPATH)
	CGO_ENABLED ?= 0
	export CGO_ENABLED
else
	VERSION := $(shell git describe --exact-match --tags 2>/dev/null)
endif

PREFIX := /usr/local
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git rev-parse --short HEAD)
GOFILES ?= $(shell git ls-files '*.go')
BUILDFLAGS ?=

ifdef GOBIN
PATH := $(GOBIN):$(PATH)
else
PATH := $(subst :,/bin:,$(shell go env GOPATH))/bin:$(PATH)
endif

LDFLAGS := -X main.commit=$(COMMIT) -X main.branch=$(BRANCH)
ifdef VERSION
	LDFLAGS += -X main.version=$(VERSION)
endif

.PHONY: telegraf
telegraf:
	go build -ldflags "$(LDFLAGS)" contrail-introspect-cli/main.go contrail-introspect-cli/command.go
