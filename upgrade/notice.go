package upgrade

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gookit/color"
	"github.com/hashicorp/go-version"

	"github.com/mszostok/version/style"
)

type (
	RenderFunc       func(in *Info) (string, error)
	PostRenderFunc   func(body string) (string, error)
	IsVerGreaterFunc func(current string, new string) bool
	Options          func(options *GitHubDetector)
)

type GitHubDetector struct {
	customRenderFn     RenderFunc
	postRenderFn       PostRenderFunc
	style              *style.Config
	repo               string
	mux                sync.RWMutex
	info               *Info
	stateFilePath      string
	isVerGreater       IsVerGreaterFunc
	updateCheckTimeout time.Duration
	recheckInterval    time.Duration
}

func NewGitHubDetector(owner, repo string, opts ...Options) *GitHubDetector {
	gh := GitHubDetector{
		style:              style.DefaultConfig(defaultLayoutGoTpl),
		stateFilePath:      "/tmp/state.yml",
		repo:               fmt.Sprintf("%s/%s", owner, repo),
		isVerGreater:       semvVerGreater,
		updateCheckTimeout: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(&gh)
	}
	//if err := options.Validate(); err != nil {
	//	return "", err
	//}

	return &gh
}

func (gh *GitHubDetector) Validate() error {
	if gh == nil {
		return errors.New("config is nil")
	}
	if gh.repo == "" {
		return errors.New("repository URL is required")
	}
	return nil
}

func (gh *GitHubDetector) Render() (string, error) {
	if err := gh.Validate(); err != nil {
		return "", err
	}
	body, err := gh.render()
	if err != nil {
		return "", err
	}

	if gh.postRenderFn != nil {
		return gh.postRenderFn(body)
	}

	return body, nil
}

func (gh *GitHubDetector) render() (string, error) {
	if gh.customRenderFn != nil {
		return gh.customRenderFn(gh.info)
	}

	renderBody := style.NewGoTemplateRender(gh.style)
	body, err := renderBody.Render(gh.info)
	if err != nil {
		return "", err
	}
	return body, nil
}

func (gh *GitHubDetector) CheckForUpdate(currentVersion string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), gh.updateCheckTimeout)
	defer cancel()

	rel, err := GetLatestRelease(ctx, gh.stateFilePath, gh.repo, gh.recheckInterval)
	if err != nil || rel == nil {
		fmt.Println(err)
		return false
	}
	if !gh.isVerGreater(currentVersion, rel.Version) {
		return false
	}

	gh.mux.Lock()
	defer gh.mux.Unlock()

	gh.info = &Info{
		Version:     currentVersion,
		NewVersion:  rel.Version,
		BrewUpgrade: fmt.Sprintf("To upgrade, run: %s", color.Gray.Sprint("brew update && brew upgrade gh")),
		ReleaseURL:  rel.URL,
	}

	return true
}

func semvVerGreater(current, new string) bool {
	current = strings.TrimPrefix(current, "v")
	new = strings.TrimPrefix(new, "v")
	currv, curre := version.NewVersion(current)
	newv, newe := version.NewVersion(new)

	return curre == nil && newe == nil && newv.GreaterThan(currv)
}
