```makefile
# The version common library import path
export VERSION_PKG=github.com/mszostok/version

# The ldflags for the go build process to set the version related data.
export GO_BUILD_VERSION_LDFLAGS=\
	-X $(VERSION_PKG).version=$(VERSION) \
	-X $(VERSION_PKG).dirtyBuild=false \
	-X $(VERSION_PKG).buildDate=$(shell date +"%Y%m%d-%T")
```
