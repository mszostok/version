package github

import (
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

type stateEntry struct {
	CheckedForUpdateAt  time.Time           `yaml:"checkedForUpdateAt"`
	ReleaseInfoResponse ReleaseInfoResponse `yaml:"releaseInfoResponse"`
}

func getStateEntry(stateFilePath string) (*stateEntry, error) {
	content, err := os.ReadFile(stateFilePath)
	switch {
	case err == nil:
	case os.IsNotExist(err):
		return nil, nil
	default:
		return nil, err
	}

	var state stateEntry
	err = yaml.Unmarshal(content, &state)
	if err != nil {
		return nil, err
	}

	state.ReleaseInfoResponse.IsFromCache = true
	return &state, nil
}

func saveStateEntry(stateFilePath string, info ReleaseInfoResponse, t time.Time) error {
	info.IsFromCache = true
	data := stateEntry{
		CheckedForUpdateAt:  t,
		ReleaseInfoResponse: info,
	}

	content, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(stateFilePath), 0o755)
	if err != nil {
		return err
	}

	err = os.WriteFile(stateFilePath, content, 0o600)
	return err
}
