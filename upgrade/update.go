package upgrade

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

// ReleaseInfoResponse stores information about a release.
type ReleaseInfoResponse struct {
	Version     string    `json:"tag_name"`
	URL         string    `json:"html_url"`
	PublishedAt time.Time `json:"published_at"`
}

// GetLatestRelease checks whether there is a newer release on GitHub.
func GetLatestRelease(ctx context.Context, stateFilePath, repo string, minRecheckTime time.Duration) (*ReleaseInfoResponse, error) {
	stateEntry, err := getStateEntry(stateFilePath)
	if err != nil {
		return nil, err
	}

	if stateEntry != nil && time.Since(stateEntry.CheckedForUpdateAt) < minRecheckTime {
		return nil, nil
	}

	releaseInfo, err := getLatestReleaseInfo(ctx, repo)
	if err != nil {
		return nil, err
	}

	err = setStateEntry(stateFilePath, time.Now())
	if err != nil {
		return nil, err
	}
	return releaseInfo, nil
}

func getLatestReleaseInfo(ctx context.Context, repo string) (*ReleaseInfoResponse, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var latestRelease ReleaseInfoResponse
	err = json.Unmarshal(rawBody, &latestRelease)
	if err != nil {
		return nil, err
	}

	return &latestRelease, nil
}

type stateEntry struct {
	CheckedForUpdateAt time.Time `yaml:"checkedForUpdateAt"`
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

	var stateEntry stateEntry
	err = yaml.Unmarshal(content, &stateEntry)
	if err != nil {
		return nil, err
	}

	return &stateEntry, nil
}

func setStateEntry(stateFilePath string, t time.Time) error {
	data := stateEntry{CheckedForUpdateAt: t}
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
