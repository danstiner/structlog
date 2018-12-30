.PHONY: build build-alpine clean test bench help default

BIN_NAME=go-structlog

GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
IMAGE_NAME := "danstiner/go-structlog"

default: test

help:
	@echo 'Management commands for go-structlog:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make get-deps        runs dep ensure, mostly used for ci.'
	
	@echo '    make clean           Clean the directory tree.'
	@echo

build:
	@echo "building ${BIN_NAME}
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X main.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X main.VersionPrerelease=DEV" -o bin/${BIN_NAME}

get-deps:
	dep ensure

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

test:
	go test ./...

bench:
	go test -bench=. ./...

profile:
	go test -gcflags=-m -cpuprofile cpu.prof -memprofile mem.prof -bench=Format10Fields github.com/danstiner/go-structlog/messagetemplates
	go tool pprof --pdf cpu.prof
	go tool pprof --pdf mem.prof
