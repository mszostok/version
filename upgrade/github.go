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
	"go.szostok.io/version/upgrade/github"
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

// LookForLatestRelease checks project's latest release. In most cases, PrintIfFoundGreater is better option.
//
// Deprecated and will be removed in 1.4.0 release, use LookForGreaterRelease instead.
//
//	Old: LookForLatestRelease(upgrade.LookForLatestReleaseInput{CurrentVersion: currentVersion})
//	New: LookForGreaterRelease(upgrade.LookForGreaterReleaseInput{CurrentVersion: currentVersion})
func (gh *GitHubDetector) LookForLatestRelease(in LookForLatestReleaseInput) (LookForLatestReleaseOutput, error) {
	out, err := gh.LookForGreaterRelease(LookForGreaterReleaseInput(in))
	return LookForLatestReleaseOutput(out), err
}

// LookForGreaterReleaseOutput holds output data for LookForGreaterRelease function.
type LookForGreaterReleaseOutput struct {
	Found       bool
	ReleaseInfo *Info
}

// LookForGreaterReleaseInput holds input data for LookForGreaterRelease function.
type LookForGreaterReleaseInput struct {
	CurrentVersion string
}

// LookForGreaterRelease checks project's latest release. In most cases, PrintIfFoundGreater is better option.
func (gh *GitHubDetector) LookForGreaterRelease(in LookForGreaterReleaseInput) (LookForGreaterReleaseOutput, error) {
	var empty LookForGreaterReleaseOutput
	if err := gh.validate(); err != nil { // TODO: move to constructor
		return empty, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), gh.updateCheckTimeout)
	defer cancel()

	rel, err := gh.getReleaseInfo(ctx)
	if err != nil || rel == nil {
		return empty, err
	}
	if !gh.isVerGreater(in.CurrentVersion, rel.Version) {
		return empty, err
	}

	return LookForGreaterReleaseOutput{
		Found: true,
		ReleaseInfo: &Info{
			Version:     in.CurrentVersion,
			NewVersion:  rel.Version,
			ReleaseURL:  rel.URL,
			IsFromCache: rel.IsFromCache,
		},
	}, nil
}

// PrintIfFoundGreater prints an upgrade notice if a newer version is available.
// It's a syntax sugar for using the LookForGreaterRelease and Render functions.
func (gh *GitHubDetector) PrintIfFoundGreater(w io.Writer, currentVersion string) error {
	rel, err := gh.LookForGreaterRelease(LookForGreaterReleaseInput{
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

func (gh *GitHubDetector) getReleaseInfo(ctx context.Context) (*github.ReleaseInfoResponse, error) {
	if gh.recheckInterval == 0 {
		return github.GetLatestRelease(ctx, gh.repo)
	}

	statePath := gh.getStateFilePath()
	return github.GetLatestReleaseWithCache(ctx, gh.repo, statePath, gh.recheckInterval)
}

func (gh *GitHubDetector) getStateFilePath() string {
	if gh.configDir != "" {
		return filepath.Join(gh.configDir, gh.stateFileName)
	}

	return filepath.Join(DefaultConfigDir(), gh.stateFileName)
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
