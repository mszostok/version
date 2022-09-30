package upgrade

// Info contains upgrade related information.
type Info struct {
	// Version represents a current CLI version.
	Version string
	// NewVersion represents the newest CLI version.
	NewVersion string
	// ReleaseURL represents the GitHub release URL.
	ReleaseURL string
	// IsFromCache if true, the external call was not executed and version was taken from state file.
	IsFromCache bool
}
