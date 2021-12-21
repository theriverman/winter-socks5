winter-socks5 [![Makefile CD](https://github.com/theriverman/winter-socks5/actions/workflows/makefile.yml/badge.svg?branch=master)](https://github.com/theriverman/winter-socks5/actions/workflows/makefile.yml)
[![GoDoc](https://godoc.org/github.com/theriverman/winter-socks5?status.svg)](https://pkg.go.dev/github.com/theriverman/winter-socks5?tab=doc "Docs @ pkg.go.dev")
=========

**WinterSOCKS5** is a CLI frontend for the following SOCKS5 Go implementations:
  * [theriverman/go-socks5](https://github.com/theriverman/go-socks5) (forked from [armon/go-socks5](https://github.com/armon/go-socks5)).
  * [txthinking/socks5](https://github.com/txthinking/socks5)
  
SOCKS (Secure Sockets) is used to route traffic between a client and a server through an intermediate proxy layer. This can be used to bypass firewalls or NATs.

See [Releases](https://github.com/theriverman/winter-socks5/releases) for pre-compiled binaries (Linux, macOS, Windows). A separate binary is provided for each backend.

# Usage
See help for instructions:
```cmd
.\socks5-cli.armon-windows-amd64.exe --help
.\socks5-cli.txthinking-windows-amd64.exe --help
```

**Note:** The name of your binary may differ from the example.

# Building
The following prerequisites must be met to build **WinterSOCKS5**:
  * go 1.17 or higher
  * Python 3.6 or higher
  * pip (Python pkg-management system)

## Building with GNU make
**Note:** Make a new git tag before building a new release version!

See the available build targets by executing the following command:
```bash
make info
```

The most common build targets are the following:
```bash
make build-armon-windows
make build-txthinking-windows
```

## Building Manually
See the contents of [Makefile](./Makefile) for reference.
