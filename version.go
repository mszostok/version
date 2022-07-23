package version

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

// Info contains versioning information.
type Info struct {
	Version    string `yaml:"version,omitempty"     json:"version,omitempty"`
	GitCommit  string `yaml:"gitCommit,omitempty"   json:"gitCommit,omitempty"`
	BuildDate  string `yaml:"buildDate,omitempty"   json:"buildDate,omitempty"`
	CommitDate string `yaml:"commitDate,omitempty"  json:"commitDate,omitempty"`
	DirtyBuild bool   `yaml:"dirtyBuild,omitempty"  json:"dirtyBuild,omitempty"`
	GoVersion  string `yaml:"goVersion,omitempty"   json:"goVersion,omitempty"`
	Compiler   string `yaml:"compiler,omitempty"    json:"compiler,omitempty"`
	Platform   string `yaml:"platform,omitempty"    json:"platform,omitempty"`

	name string
}

// Get returns the overall codebase version.
// It's for detecting what code a binary was built from.
//
// Version related variables are resolved in such order:
//   1. Go -ldflags
//   2. or debug.ReadBuildInfo() in Go 1.18+
//      * version is set only if the binary is built with "go install url/tool@version".
//      * commit is taken from the vcs.revision tag.
//      * commitDate is taken from the vcs.time tag.
//      * dirtyBuild is taken from the vcs.modified tag.
//   3. in their absence fallback to the settings in ./base.go.
func Get(name ...string) Info {
	var n string
	if len(name) > 0 {
		n = name[0]
	}
	return Info{
		name:       n,
		Version:    version,
		GitCommit:  commit,
		BuildDate:  buildDate,
		CommitDate: commitDate,
		DirtyBuild: dirtyBuild,
		GoVersion:  runtime.Version(),
		Compiler:   runtime.Compiler,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// CollectFromBuildInfo tries to set the build information embedded in the running binary via Go module.
// It doesn't override data if were already set by Go -ldflags.
func CollectFromBuildInfo() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	if version == unknownVersion {
		version = info.Main.Version
	}

	for _, kv := range info.Settings {
		switch kv.Key {
		case "vcs.revision":
			if commit == unknownProperty {
				commit = kv.Value
			}
		case "vcs.time":
			if commitDate == unknownProperty {
				commitDate = kv.Value
			}
		case "vcs.modified":
			dirtyBuild = kv.Value == "true"
		}
	}
}
