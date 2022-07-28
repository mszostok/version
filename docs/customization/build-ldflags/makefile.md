# Makefile

This is an example `Makefile` to build your Go application with `go.szostok.io/version` support:

```makefile
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)
VERSION?="dev"

# The ldflags for the go build process to set the version related data.
GO_BUILD_VERSION_LDFLAGS=\
  -X go.szostok.io/version.version=$(VERSION)\
  -X go.szostok.io/version.buildDate=$(shell date +"%Y-%m-%dT%H:%M:%S%z")

build:
		GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags="$(GO_BUILD_LDFLAGS)" -o example ./example/
.PHONY: build
```

The rest properties are set based on the built-in data. However, if you want to have a full customization, use:

```makefile
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)
VERSION?="dev"

# The ldflags for the go build process to set the version related data.
GO_BUILD_VERSION_LDFLAGS=\
  -X go.szostok.io/version.version=$(VERSION) \
  -X go.szostok.io/version.buildDate=$(shell date +"%Y-%m-%dT%H:%M:%S%z") \
  -X go.szostok.io/commit=$(shell git rev-parse --short HEAD) \
  -X go.szostok.io/commitDate=$(shell git log -1 --date=format:"%Y-%m-%dT%H:%M:%S%z" --format=%cd) \
  -X go.szostok.io/dirtyBuild=false

build:
		GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags="$(GO_BUILD_LDFLAGS)" -o example ./example/
.PHONY: build
```
