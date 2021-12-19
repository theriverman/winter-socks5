# BINARY SUFFIX (for proper Windows support)
ifeq ($(OS),Windows_NT)
	BINARY_SUFFIX := .exe
else
	BINARY_SUFFIX :=
endif

# BUILD-TIME VARIABLES
OS_NAME := $(shell uname -s | tr A-Z a-z)
BUILD_TYPE := released
CURRENT_TIME := $(shell date '+%c')
LATEST_GIT_TAG := $(shell git describe --tags --abbrev=0)
LATEST_GIT_COMMIT := $(shell git log -n 1 --pretty=format:"%H")

# GLOBALS
BINARY_NAME := socks5-cli
GO_LD_FLAGS := -ldflags "-X 'main.app_built_date=$(CURRENT_TIME)' -X 'main.app_build_type=$(BUILD_TYPE)' -X 'main.app_sem_version=$(LATEST_GIT_TAG)' -X 'main.git_commit=$(LATEST_GIT_COMMIT)'"
BACKEND_armon := -ldflags "-X 'main.proxyBackend=github.com/armon/go-socks5'"
BACKEND_txthinking := -ldflags "-X 'main.proxyBackend=github.com/txthinking/socks5'"

# monkey-patch for Linux => Windows cross-compilation
GO_BIN_DIR := $(shell go env GOPATH)/bin
export PATH := $(GO_BIN_DIR):$(PATH)

info:
	@echo "Choose from the following targets:"
	@echo "  * build-armon (defaults to your host/os settings)"
	@echo "  * build-armon-darwin"
	@echo "  * build-armon-linux"
	@echo "  * build-armon-windows"
	@echo "    -----------------------------------------------"
	@echo "  * build-txthinking (defaults to your host/os settings)"
	@echo "  * build-txthinking-darwin"
	@echo "  * build-txthinking-linux"
	@echo "  * build-txthinking-windows"

build-armon:
	@echo "Building with backend: github.com/armon/go-socks5 for $(shell go env GOOS)/$(shell go env GOARCH)"
	@go build $(GO_LD_FLAGS) $(BACKEND_armon) --tags armon -o dist/$(BINARY_NAME).armon$(BINARY_SUFFIX)

build-armon-darwin:
	@go build $(GO_LD_FLAGS) --tags armon -o dist/$(BINARY_NAME).armon$(BINARY_SUFFIX)
	@GOOS=darwin GOARCH=amd64 go build $(GO_LD_FLAGS) $(BACKEND_armon) --tags armon -o dist/$(BINARY_NAME).armon-darwin-amd64
	@GOOS=darwin GOARCH=arm64 go build $(GO_LD_FLAGS) $(BACKEND_armon) --tags armon -o dist/$(BINARY_NAME).armon-darwin-arm64

build-armon-linux:
	@GOOS=linux GOARCH=386 go build $(GO_LD_FLAGS) $(BACKEND_armon) --tags armon -o dist/$(BINARY_NAME).armon-linux-386
	@GOOS=linux GOARCH=amd64 go build $(GO_LD_FLAGS) $(BACKEND_armon) --tags armon -o dist/$(BINARY_NAME).armon-linux-amd64
	@GOOS=linux GOARCH=arm go build $(GO_LD_FLAGS) $(BACKEND_armon) --tags armon -o dist/$(BINARY_NAME).armon-linux-arm
	@GOOS=linux GOARCH=arm64 go build $(GO_LD_FLAGS) $(BACKEND_armon) --tags armon -o dist/$(BINARY_NAME).armon-linux-arm64

build-armon-windows:
	@python versioninfo.py $(LATEST_GIT_TAG)
	@GOOS=windows go generate
	@GOOS=windows GOARCH=386 go build $(GO_LD_FLAGS) $(BACKEND_armon) --tags armon -o dist/$(BINARY_NAME).armon-windows-386.exe
	@GOOS=windows GOARCH=amd64 go build $(GO_LD_FLAGS) $(BACKEND_armon) --tags armon -o dist/$(BINARY_NAME).armon-windows-amd64.exe

build-txthinking:
	@echo "Building with backend: github.com/txthinking/socks5 for $(shell go env GOOS)/$(shell go env GOARCH)"
	@go build $(GO_LD_FLAGS) $(BACKEND_txthinking) --tags txthinking -o dist/$(BINARY_NAME).txthinking$(BINARY_SUFFIX)

build-txthinking-darwin:
	@go build $(GO_LD_FLAGS) --tags txthinking -o dist/$(BINARY_NAME).txthinking$(BINARY_SUFFIX)
	@GOOS=darwin GOARCH=amd64 go build $(GO_LD_FLAGS) $(BACKEND_txthinking) --tags txthinking -o dist/$(BINARY_NAME).txthinking-darwin-amd64
	@GOOS=darwin GOARCH=arm64 go build $(GO_LD_FLAGS) $(BACKEND_txthinking) --tags txthinking -o dist/$(BINARY_NAME).txthinking-darwin-arm64

build-txthinking-linux:
	@GOOS=linux GOARCH=386 go build $(GO_LD_FLAGS) $(BACKEND_txthinking) --tags txthinking -o dist/$(BINARY_NAME).txthinking-linux-386
	@GOOS=linux GOARCH=amd64 go build $(GO_LD_FLAGS) $(BACKEND_txthinking) --tags txthinking -o dist/$(BINARY_NAME).txthinking-linux-amd64
	@GOOS=linux GOARCH=arm go build $(GO_LD_FLAGS) $(BACKEND_txthinking) --tags txthinking -o dist/$(BINARY_NAME).txthinking-linux-arm
	@GOOS=linux GOARCH=arm64 go build $(GO_LD_FLAGS) $(BACKEND_txthinking) --tags txthinking -o dist/$(BINARY_NAME).txthinking-linux-arm64

build-txthinking-windows:
	@python versioninfo.py $(LATEST_GIT_TAG)
	@GOOS=windows go generate
	@GOOS=windows GOARCH=386 go build $(GO_LD_FLAGS) $(BACKEND_txthinking) --tags txthinking -o dist/$(BINARY_NAME).txthinking-windows-386.exe
	@GOOS=windows GOARCH=amd64 go build $(GO_LD_FLAGS) $(BACKEND_txthinking) --tags txthinking -o dist/$(BINARY_NAME).txthinking-windows-amd64.exe

check: install-dev-dependencies clean build

clean:
	rm -rf ./dist/*
	rm -rf ./versioninfo.json
	rm -rf ./resource_*.syso

install-dev-dependencies:
	go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest
	pip install semver
