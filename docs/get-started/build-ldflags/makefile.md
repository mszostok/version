# Makefile

This is an example `Makefile` to build your Go application and override the default version information set by the `go.szostok.io/version` package.

```makefile
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)
VERSION?="dev"

# The ldflags for the Go build process to set the version related data
GO_BUILD_VERSION_LDFLAGS=\
  -X go.szostok.io/version.version=$(VERSION)\
  -X go.szostok.io/version.buildDate=$(shell date +"%Y-%m-%dT%H:%M:%S%z")

build:
		GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags="$(GO_BUILD_LDFLAGS)" -o example ./example/
.PHONY: build
```

The remaining properties are set based on the built-in data. However, for full customization, use:

```makefile
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)
VERSION?="dev"

# The ldflags for the Go build process to set the version related data
GO_BUILD_VERSION_LDFLAGS=\
  -X go.szostok.io/version.version=$(VERSION) \
  -X go.szostok.io/version.buildDate=$(shell date +"%Y-%m-%dT%H:%M:%S%z") \
  -X go.szostok.io/version.commit=$(shell git rev-parse --short HEAD) \
  -X go.szostok.io/version.commitDate=$(shell git log -1 --date=format:"%Y-%m-%dT%H:%M:%S%z" --format=%cd) \
  -X go.szostok.io/version.dirtyBuild=false

build:
		GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags="$(GO_BUILD_LDFLAGS)" -o example ./example/
.PHONY: build
```
