package upgrade

import (
	"time"

	"github.com/mszostok/version/style"
)

var noop = func(options *GitHubDetector) {}

// WithRenderer sets a custom renderer function.
func WithRenderer(renderer RenderFunc) Options {
	return func(options *GitHubDetector) {
		options.customRenderFn = renderer
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

// WithBoxed sets a box style.
// Disabled by default.
func WithBoxed(color BoxColor) Options {
	return func(options *GitHubDetector) {
		options.style.Layout.GoTemplate = forBoxLayoutGoTpl
		options.boxed = true
		options.boxedColor = color
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
// Defaults to 24h.
func WithMinElapseTimeForRecheck(interval time.Duration) Options {
	return func(options *GitHubDetector) {
		options.recheckInterval = interval
	}
}
