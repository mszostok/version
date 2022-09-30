package upgrade

import (
	"context"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"time"

	semver "github.com/hashicorp/go-version"

	"go.szostok.io/version/style"
	"go.szostok.io/version/term"
)

var defaultLayoutGoTpl = `
  │ A new release is available: {{ .Version }} → {{ .NewVersion | Green }}
  │ {{ .ReleaseURL  | Underline | Blue }}
`

// GitHubDetector provides functionality to check GitHub for project's latest release.
type GitHubDetector struct {
	customRenderFn     RenderFunc
	postRenderFn       PostRenderFunc
	style              *style.Config
	repo               string
	stateFileName      string
	isVerGreater       IsVerGreaterFunc
	updateCheckTimeout time.Duration
	recheckInterval    time.Duration
	configDir          string
}

// NewGitHubDetector returns GitHubDetector instance.
func NewGitHubDetector(owner, repo string, opts ...Options) *GitHubDetector {
	ghRepo := fmt.Sprintf("%s/%s", owner, repo)
	gh := GitHubDetector{
		style:              style.DefaultConfig(defaultLayoutGoTpl),
		stateFileName:      ghRepo,
		repo:               ghRepo,
		isVerGreater:       semvVerGreater,
		updateCheckTimeout: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(&gh)
	}

	return &gh
}

// Render returns rendered input version with configured style.
func (gh *GitHubDetector) Render(info *Info, isSmartTerminal bool) (string, error) {
	body, err := gh.render(info, isSmartTerminal)
	if err != nil {
		return "", err
	}

	if gh.postRenderFn != nil {
		return gh.postRenderFn(body, isSmartTerminal)
	}

	return body, nil
}

// LookForLatestReleaseOutput holds output data for LookForLatestRelease function.
type LookForLatestReleaseOutput struct {
	Found       bool
	ReleaseInfo *Info
}

// LookForLatestReleaseInput holds input data for LookForLatestRelease function.
type LookForLatestReleaseInput struct {
	CurrentVersion string
}

// LookForLatestRelease if a given time elapsed, check project's latest release.
func (gh *GitHubDetector) LookForLatestRelease(in LookForLatestReleaseInput) (LookForLatestReleaseOutput, error) {
	var empty LookForLatestReleaseOutput
	if err := gh.validate(); err != nil { // TODO: move to constructor
		return empty, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), gh.updateCheckTimeout)
	defer cancel()

	statePath := gh.getStateFilePath()

	rel, err := GetLatestRelease(ctx, statePath, gh.repo, gh.recheckInterval)
	if err != nil || rel == nil {
		return empty, err
	}
	if !gh.isVerGreater(in.CurrentVersion, rel.Version) {
		return empty, err
	}

	return LookForLatestReleaseOutput{
		Found: true,
		ReleaseInfo: &Info{
			Version:     in.CurrentVersion,
			NewVersion:  rel.Version,
			ReleaseURL:  rel.URL,
			IsFromCache: rel.IsFromCache,
		},
	}, nil
}

func (gh *GitHubDetector) getStateFilePath() string {
	if gh.configDir != "" {
		return filepath.Join(gh.configDir, gh.stateFileName)
	}

	return filepath.Join(DefaultConfigDir(), gh.stateFileName)
}

// PrintIfFoundGreater prints an upgrade notice if a newer version is available.
// It's a syntax sugar for using the LookForLatestRelease and Render functions.
func (gh *GitHubDetector) PrintIfFoundGreater(w io.Writer, currentVersion string) error {
	rel, err := gh.LookForLatestRelease(LookForLatestReleaseInput{
		CurrentVersion: currentVersion,
	})
	if err != nil {
		return err
	}

	if !rel.Found {
		return nil
	}
	out, err := gh.Render(rel.ReleaseInfo, term.IsSmart(w))
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(w, out)
	return err
}

func (gh *GitHubDetector) render(info *Info, isSmartTerminal bool) (string, error) {
	if gh.customRenderFn != nil {
		return gh.customRenderFn(info, isSmartTerminal)
	}

	renderBody := style.NewGoTemplateRender(gh.style)
	body, err := renderBody.Render(info, isSmartTerminal)
	if err != nil {
		return "", err
	}
	return body, nil
}

func (gh *GitHubDetector) validate() error {
	if gh == nil {
		return errors.New("config cannot be nil")
	}
	if gh.repo == "" {
		return errors.New("repository URL is required")
	}
	return nil
}

func semvVerGreater(current, got string) bool {
	current = strings.TrimPrefix(current, "v")
	got = strings.TrimPrefix(got, "v")
	currv, curre := semver.NewVersion(current)
	newv, newe := semver.NewVersion(got)

	return curre == nil && newe == nil && newv.GreaterThan(currv)
}
