package upgrade

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// ReleaseInfoResponse stores information about a release.
type ReleaseInfoResponse struct {
	IsFromCache bool   `yaml:"cached"`
	Version     string `yaml:"version"`
	URL         string `yaml:"URL"`
}

// GetLatestRelease checks whether there is a newer release on GitHub. If yes, returns it, otherwise returns nil.
func GetLatestRelease(ctx context.Context, stateFilePath, repo string, minRecheckTime time.Duration) (*ReleaseInfoResponse, error) {
	stateEntry, err := getStateEntry(stateFilePath)
	if err != nil {
		return nil, err
	}

	if stateEntry != nil && time.Since(stateEntry.CheckedForUpdateAt) < minRecheckTime {
		return &stateEntry.ReleaseInfoResponse, nil
	}

	releaseInfo, err := getLatestReleaseInfo(ctx, repo)
	if err != nil {
		return nil, err
	}

	err = saveStateEntry(stateFilePath, *releaseInfo, time.Now())
	if err != nil {
		return nil, err
	}
	return releaseInfo, nil
}

func getLatestReleaseInfo(ctx context.Context, repo string) (*ReleaseInfoResponse, error) {
	url := fmt.Sprintf("https://github.com/%s/releases/latest", repo)

	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	client := newHTTPClient()
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusFound {
		return nil, fmt.Errorf("incorrect status code: %d", res.StatusCode)
	}

	loc := res.Header.Get("Location")
	if loc == "" {
		return nil, fmt.Errorf("unable to determine release version")
	}
	version := loc[strings.LastIndex(loc, "/")+1:]
	return &ReleaseInfoResponse{
		Version: version,
		URL:     loc,
	}, nil
}

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
