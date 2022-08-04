package upgrade

import "time"

// Info contains upgrade related information.
type Info struct {
	// Version represents a current CLI version.
	Version string
	// NewVersion represents the newest CLI version.
	NewVersion string
	// ReleaseURL represents the GitHub release URL.
	ReleaseURL string
	// PublishedAt represents the release publish date.
	PublishedAt time.Time
}
