package upgrade

import "time"

// Info contains upgrade related information.
type Info struct {
	Version     string
	NewVersion  string
	ReleaseURL  string
	PublishedAt time.Time
}
