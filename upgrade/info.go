package upgrade

import "time"

type Info struct {
	Version     string
	NewVersion  string
	BrewUpgrade string
	ReleaseURL  string
	PublishedAt time.Time
}
