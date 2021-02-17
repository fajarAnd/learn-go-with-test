.PHONY: default help build package tag push helm helm-migration run test clean

SHELL         = /bin/bash
APP_NAME      = learn-go-with-test
VERSION      := $(shell git describe --always --tags)
GIT_COMMIT    = $(shell git rev-parse HEAD)
BUILD_DATE    = $(shell date '+%Y-%m-%d-%H:%M:%S')

default: help

test:
	@echo "Testing ${APP_NAME}"
	go get github.com/mattn/goveralls && go test ./...

coverage:
	@echo "Checking unit test coverage ${APP_NAME} ${VERSION}"
	go get github.com/mattn/goveralls && \
		go test -cover -coverprofile=coverage.out ./... && \
		goveralls -coverprofile=coverage.out -repotoken=${COVERALLS_REPO_TOKEN_LEARN_GO_WITH_TEST} && \
		exit 0 || exit 1
