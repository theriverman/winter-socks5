go-socks5-cli [![Makefile CD](https://github.com/theriverman/go-socks5-cli/actions/workflows/makefile.yml/badge.svg?branch=master)](https://github.com/theriverman/go-socks5-cli/actions/workflows/makefile.yml)
[![GoDoc](https://godoc.org/github.com/theriverman/go-socks5-cli?status.svg)](https://pkg.go.dev/github.com/theriverman/go-socks5-cli?tab=doc "Docs @ pkg.go.dev")
=========

This **go-socks5-cli** application is a CLI wrapper around `theriverman/go-socks5` (forked from `armon/go-socks5`). The **go-socks5-cli** application uses the `socks5` package from `go-socks5` which implements a [SOCKS5 server](http://en.wikipedia.org/wiki/SOCKS).

SOCKS (Secure Sockets) is used to route traffic between a client and a server through an intermediate proxy layer. This can be used to bypass firewalls or NATs.

# Usage
```bash
$ socks5-cli.exe server
# or
$ socks5-cli.exe server --address 0.0.0.0 --port 1080
```
**Note:** The name of your binary may differ from the example.

# Building with GNU make
**Note:** Make a new git tag before building a new release version!

```bash
make build
```

# Building Manually
**Note:** Make a new git tag before building a new release version!

Store some pre-build information:
```bash
export BINARY_NAME="socks5-cli"
export BINARY_SUFFIX=".exe"
CURRENT_TIME=$(date '+%c')
LATEST_GIT_TAG=$(git describe --tags --abbrev=0)
LATEST_GIT_COMMIT=$(git log -n 1 --pretty=format:"%H")
export CURRENT_TIME
export LATEST_GIT_TAG
export LATEST_GIT_COMMIT
```

Build the binary by executing the following commands:
```bash
go generate
go build -ldflags "-X 'main.app_built_date=$CURRENT_TIME' -X 'main.app_build_type=released' -X 'main.app_sem_version=$LATEST_GIT_TAG' -X 'main.git_commit=$LATEST_GIT_COMMIT'" -o dist/$BINARY_NAME$BINARY_SUFFIX
```
