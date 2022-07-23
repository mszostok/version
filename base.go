package version

const (
	unknownVersion  = "(devel)"
	unknownProperty = "N/A"
)

// Fallback data used when versioning information is not provided.
var (
	version    = unknownVersion
	commit     = unknownProperty
	buildDate  = unknownProperty
	commitDate = unknownProperty
	dirtyBuild = false
)
