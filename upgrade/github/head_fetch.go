package github

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// The previous version used the GitHub API url (see: https://github.com/mszostok/version/blob/v1.0.0/upgrade/update.go#L47)
// But with the official API it's easy to hit rate limiting.
// TODO: Alternative strategy that can be implemented is based on the E-Tag. As a result we will cache API call and also have full release info.
const gitHubLatestReleaseURLFormat = "https://github.com/%s/releases/latest"

// ReleaseInfoResponse stores information about a release.
type ReleaseInfoResponse struct {
	IsFromCache bool   `yaml:"cached"`
	Version     string `yaml:"version"`
	URL         string `yaml:"URL"`
}

// GetLatestRelease returns the latest release found on GitHub.
// The latest release is the most recent non-prerelease, non-draft release, sorted by the created_at attribute.
// The created_at attribute is the date of the commit used for the release, and not the date when the release was drafted or published.
func GetLatestRelease(ctx context.Context, repo string) (*ReleaseInfoResponse, error) {
	url := fmt.Sprintf(gitHubLatestReleaseURLFormat, repo)

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

// GetLatestReleaseWithCache returns the latest release found on GitHub. If minRecheckTime not elapsed, cached version is returned.
// The latest release is the most recent non-prerelease, non-draft release, sorted by the created_at attribute.
// The created_at attribute is the date of the commit used for the release, and not the date when the release was drafted or published.
func GetLatestReleaseWithCache(ctx context.Context, repo, stateFilePath string, minRecheckTime time.Duration) (*ReleaseInfoResponse, error) {
	stateEntry, err := getStateEntry(stateFilePath)
	if err != nil {
		return nil, err
	}

	if stateEntry != nil && time.Since(stateEntry.CheckedForUpdateAt) < minRecheckTime {
		return &stateEntry.ReleaseInfoResponse, nil
	}

	releaseInfo, err := GetLatestRelease(ctx, repo)
	if err != nil {
		return nil, err
	}

	err = saveStateEntry(stateFilePath, *releaseInfo, time.Now())
	if err != nil {
		return nil, err
	}
	return releaseInfo, nil
}
