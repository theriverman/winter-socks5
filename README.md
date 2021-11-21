# Introduction
N/A

# Usage
```bash
$ ./dist/socks5-cli.exe server
# or
$ ./dist/socks5-cli.exe server --address 0.0.0.0 --port 1080
```

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
