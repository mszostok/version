package upgrade

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/hashicorp/go-version"

	"github.com/mszostok/version/style"
)

const stateFileName = "upgrade-state.yaml"

type (
	// RenderFunc represents render function signature.
	RenderFunc func(in *Info) (string, error)
	// PostRenderFunc represents post render function signature.
	PostRenderFunc func(body string) (string, error)
	// IsVerGreaterFunc represents version check function signature.
	IsVerGreaterFunc func(current string, new string) bool
	// Options represents function mutating default options.
	Options func(options *GitHubDetector)
)

// GitHubDetector provides functionality to check GitHub for project's latest release.
type GitHubDetector struct {
	customRenderFn     RenderFunc
	postRenderFn       PostRenderFunc
	style              *style.Config
	repo               string
	stateFilePath      string
	isVerGreater       IsVerGreaterFunc
	updateCheckTimeout time.Duration
	recheckInterval    time.Duration
}

// NewGitHubDetector returns GitHubDetector instance.
func NewGitHubDetector(owner, repo string, opts ...Options) *GitHubDetector {
	gh := GitHubDetector{
		style:              style.DefaultConfig(defaultLayoutGoTpl),
		stateFilePath:      filepath.Join(ConfigDir(), stateFileName),
		repo:               fmt.Sprintf("%s/%s", owner, repo),
		isVerGreater:       semvVerGreater,
		updateCheckTimeout: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(&gh)
	}

	return &gh
}

// Render returns rendered input version with configured style.
func (gh *GitHubDetector) Render(info *Info) (string, error) {
	body, err := gh.render(info)
	if err != nil {
		return "", err
	}

	if gh.postRenderFn != nil {
		return gh.postRenderFn(body)
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

	rel, err := GetLatestRelease(ctx, gh.stateFilePath, gh.repo, gh.recheckInterval)
	if err != nil || rel == nil {
		return empty, err
	}
	if !gh.isVerGreater(in.CurrentVersion, rel.Version) {
		return empty, err
	}

	return LookForLatestReleaseOutput{
		Found: true,
		ReleaseInfo: &Info{
			Version:    in.CurrentVersion,
			NewVersion: rel.Version,
			ReleaseURL: rel.URL,
		},
	}, nil
}

func (gh *GitHubDetector) render(info *Info) (string, error) {
	if gh.customRenderFn != nil {
		return gh.customRenderFn(info)
	}

	renderBody := style.NewGoTemplateRender(gh.style)
	body, err := renderBody.Render(info)
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
	currv, curre := version.NewVersion(current)
	newv, newe := version.NewVersion(got)

	return curre == nil && newe == nil && newv.GreaterThan(currv)
}
