

TAGS=netgo osusergo sqlite_omit_load_extension
VERSION=$(shell git describe --tags --abbrev=0 | cut -c 2-)
COMMIT=$(shell git rev-parse --verify HEAD)
DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LD_FLAGS=-s -w -linkmode external -extldflags "-static" -X main.BuildDate=$(DATE) -X main.BuildMode=prod -X main.BuildCommit=$(COMMIT) -X main.BuildVersion=$(VERSION)
BUILD_DIR=./build
PWD=$(shell pwd)
GOLANG_CROSS_VERSION=v1.22.0

license-dir:
	mkdir -p build/license || true

download-tools:
	go install github.com/Khan/genqlient
	go install golang.org/x/tools/cmd/goimports@v0.1.10

generate-go:
	genqlient

generate: generate-go

lint-go:
	go vet ./...
	goimports -l $(shell find . -type f -name '*.go' -not -path "./graph/*")

lint: lint-go

format-go:
	goimports -w $(shell find . -type f -name '*.go' -not -path "./graph/*")
	gofmt -w $(shell find . -type f -name '*.go' -not -path "./graph/*")

format: format-go

run: format-go
	go run ./main.go -i http://localhost:3030 -u admin -p admin -t http://localhost:3031 -u admin -p admin

test-go:
	go test --race -coverprofile=coverage.txt -covermode=atomic ./...

test-js:
	(cd ui && CI=true yarn test)

test: test-go test-js

install-go:
	go mod download

install: install-go

