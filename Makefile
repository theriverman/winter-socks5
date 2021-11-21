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

build:
	go generate
	go build $(GO_LD_FLAGS) -o dist/$(BINARY_NAME)$(BINARY_SUFFIX)

check: clean build

clean:
	rm -rf ./dist

compile-all:
	echo "Compiling for every OS and Platform"
	go generate
	GOOS=darwin GOARCH=amd64 go build $(GO_LD_FLAGS) -o dist/$(BINARY_NAME)-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build $(GO_LD_FLAGS) -o dist/$(BINARY_NAME)-darwin-arm64
	GOOS=linux GOARCH=386 go build $(GO_LD_FLAGS) -o dist/$(BINARY_NAME)-linux-386
	GOOS=linux GOARCH=amd64 go build $(GO_LD_FLAGS) -o dist/$(BINARY_NAME)-linux-amd64
	GOOS=linux GOARCH=arm go build $(GO_LD_FLAGS) -o dist/$(BINARY_NAME)-linux-arm
	GOOS=linux GOARCH=arm64 go build $(GO_LD_FLAGS) -o dist/$(BINARY_NAME)-linux-arm64
	GOOS=windows GOARCH=386 go build $(GO_LD_FLAGS) -o dist/$(BINARY_NAME)-windows-386.exe
	GOOS=windows GOARCH=amd64 go build $(GO_LD_FLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe

run:
	go run main.go
