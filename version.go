package version

import (
	"encoding/json"
	"fmt"
	"runtime"
	"runtime/debug"
	"strconv"
	"sync"

	"gopkg.in/yaml.v3"
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

	Meta        Meta        `json:"-" yaml:"-"`
	ExtraFields ExtraFields `json:"-" yaml:"-"`
}

// MarshalJSON marshal data into JONS but ensures that extra fields are inlined.
func (p *Info) MarshalJSON() ([]byte, error) {
	if p == nil {
		return nil, nil
	}

	type out map[string]interface{}
	type Alias Info
	f, err := json.Marshal(p.ExtraFields)
	if err != nil {
		return nil, fmt.Errorf("while marshaling extra fields: %w", err)
	}
	i, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, fmt.Errorf("while marshaling info fields: %w", err)
	}

	var ff out
	err = json.Unmarshal(f, &ff)
	if err != nil {
		return nil, fmt.Errorf("while unmarsahling extra fields: %w", err)
	}
	err = json.Unmarshal(i, &ff)
	if err != nil {
		return nil, fmt.Errorf("while unmarsahling info fields: %w", err)
	}

	return json.Marshal(ff)
}

// MarshalYAML marshal data into YAML but ensures that extra fields are inlined.
func (p *Info) MarshalYAML() (interface{}, error) {
	if p == nil {
		return nil, nil
	}

	type (
		InlinedExtra map[string]interface{}
		Alias        Info
		Output       struct {
			Alias        `yaml:",inline"`
			InlinedExtra `yaml:",omitempty,inline"`
		}
	)

	extra, err := yaml.Marshal(p.ExtraFields)
	if err != nil {
		return nil, fmt.Errorf("while marshaling extra fields: %w", err)
	}

	var mapExtra InlinedExtra
	err = yaml.Unmarshal(extra, &mapExtra)
	if err != nil {
		return nil, fmt.Errorf("while unmarsahling extra fields: %w", err)
	}

	return &Output{
		Alias:        Alias(*p),
		InlinedExtra: mapExtra,
	}, nil
}

type (
	// Meta holds additional version metadata.
	Meta struct {
		CLIName string
	}

	// ExtraFields provides an option to pass extra fields.
	ExtraFields any
)

var collectOnce sync.Once

// Get returns the overall codebase version.
// It's for detecting what code a binary was built from.
//
// Version related variables are resolved in such order:
//  1. Go -ldflags
//  2. or debug.ReadBuildInfo() in Go 1.18+
//     * version is set only if the binary is built with "go install url/tool@version".
//     * commit is taken from the vcs.revision tag.
//     * commitDate is taken from the vcs.time tag.
//     * dirtyBuild is taken from the vcs.modified tag.
//  3. in their absence fallback to the settings in ./base.go.
func Get() *Info {
	collectOnce.Do(collectFromBuildInfo)
	dirty, _ := strconv.ParseBool(dirtyBuild)
	return &Info{
		Meta:       Meta{CLIName: name},
		Version:    version,
		GitCommit:  commit,
		BuildDate:  buildDate,
		CommitDate: commitDate,
		DirtyBuild: dirty,
		GoVersion:  runtime.Version(),
		Compiler:   runtime.Compiler,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// collectFromBuildInfo tries to set the build information embedded in the running binary via Go module.
// It doesn't override data if were already set by Go -ldflags.
func collectFromBuildInfo() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	if version == unknownVersion && info.Main.Version != "" {
		version = info.Main.Version
	}

	for _, kv := range info.Settings {
		switch kv.Key {
		case "vcs.revision":
			if commit == unknownProperty && kv.Value != "" {
				commit = kv.Value
			}
		case "vcs.time":
			if commitDate == unknownProperty && kv.Value != "" {
				commitDate = kv.Value
			}
		case "vcs.modified":
			if dirtyBuild == unknownProperty && kv.Value != "" {
				dirtyBuild = kv.Value
			}
		}
	}
}
