package version

import "os"

const (
	unknownVersion  = "(devel)"
	unknownProperty = "N/A"
)

// Fallback data used when versioning information is not provided.
var (
	name       = os.Args[0]
	version    = unknownVersion
	commit     = unknownProperty
	buildDate  = unknownProperty
	commitDate = unknownProperty
	dirtyBuild = unknownProperty
)
