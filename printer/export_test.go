package printer

import "go.szostok.io/version/upgrade"

// SetOutputFormat is used only for testing purposes. It's in _test.go so it won't be available as public API.
func (r *Container) SetOutputFormat(in OutputFormat) {
	r.output = in
}

// GetUpgradeNotice is used only for testing purposes. It's in _test.go so it won't be available as public API.
func (r *Container) GetUpgradeNotice() *upgrade.GitHubDetector {
	return r.upgradeNotice
}
