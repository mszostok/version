package upgrade

import (
	"time"

	"go.szostok.io/version/style"
)

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

var noop = func(options *GitHubDetector) {}

// WithRenderer sets a custom renderer function.
func WithRenderer(renderer RenderFunc) Options {
	return func(options *GitHubDetector) {
		options.customRenderFn = renderer
	}
}

// WithPostRenderHook sets a custom post render function.
func WithPostRenderHook(renderer PostRenderFunc) Options {
	return func(options *GitHubDetector) {
		options.postRenderFn = renderer
	}
}

// WithFormatting sets a custom pretty formatting.
func WithFormatting(formatting *style.Formatting) Options {
	if formatting == nil {
		return noop
	}

	return func(options *GitHubDetector) {
		options.style.Formatting = *formatting
	}
}

// WithLayout sets a custom pretty layout.
func WithLayout(layout *style.Layout) Options {
	if layout == nil {
		return noop
	}

	return func(options *GitHubDetector) {
		options.style.Layout = *layout
	}
}

// WithStyle sets a custom pretty style.
func WithStyle(cfg *style.Config) Options {
	if cfg == nil {
		return noop
	}
	return func(options *GitHubDetector) {
		options.style = cfg
	}
}

// WithUpdateCheckTimeout sets max duration time for update check operation.
// Defaults to 10s.
func WithUpdateCheckTimeout(timeout time.Duration) Options {
	return func(options *GitHubDetector) {
		options.updateCheckTimeout = timeout
	}
}

// WithIsVersionGreater sets a custom function to compare release versions.
// Defaults to a SemVer check.
func WithIsVersionGreater(comparator IsVerGreaterFunc) Options {
	return func(options *GitHubDetector) {
		options.isVerGreater = comparator
	}
}

// WithMinElapseTimeForRecheck sets the minimum time that must elapse before checking for a new release.
// Defaults to 0.
func WithMinElapseTimeForRecheck(interval time.Duration) Options {
	return func(options *GitHubDetector) {
		options.recheckInterval = interval
	}
}
