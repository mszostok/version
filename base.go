package version

import "strings"

const (
	unknownVersion  = "(devel)"
	unknownProperty = "N/A"
)

// Fallback data used when versioning information is not provided.
var (
	name       = unknownProperty
	version    = unknownVersion
	commit     = unknownProperty
	buildDate  = unknownProperty
	commitDate = unknownProperty
	dirtyBuild = unknownProperty
)

// A Struct for Version Fields
type Field uint16

const (
	FieldMeta Field = 1 << iota
	FieldVersion
	FieldGitCommit
	FieldBuildDate
	FieldCommitDate
	FieldDirtyBuild
	FieldGoVersion
	FieldCompiler
	FieldPlatform
	FieldExtraFields

	typeRevMaxKey
)

// common presets
const (
	AllFields       Field = GoRuntimeFields | VCSFields | FieldMeta | FieldVersion | FieldExtraFields
	GoRuntimeFields       = FieldGoVersion | FieldCompiler | FieldPlatform
	VCSFields             = FieldGitCommit | FieldBuildDate | FieldCommitDate | FieldDirtyBuild
)

// Returns the string value for a Field struct
func (f Field) String() string {
	if f >= typeRevMaxKey {
		return "unknown"
	}

	switch f {
	case FieldVersion:
		return "Version"
	case FieldGitCommit:
		return "GitCommit"
	case FieldBuildDate:
		return "BuildDate"
	case FieldCommitDate:
		return "CommitDate"
	case FieldDirtyBuild:
		return "DirtyBuild"
	case FieldGoVersion:
		return "GoVersion"
	case FieldCompiler:
		return "Compiler"
	case FieldPlatform:
		return "Platform"
	case FieldMeta:
		return "Meta"
	case FieldExtraFields:
		return "ExtraFields"
	case AllFields:
		return "AllFields"
	case GoRuntimeFields:
		return "GoRuntimeFields"
	case VCSFields:
		return "VCSFields"
	}

	// Print Multiple Fields
	var fields []string
	for key := FieldMeta; key < typeRevMaxKey; key <<= 1 {
		if f&key != 0 {
			fields = append(fields, key.String())
		}
	}
	return strings.Join(fields, " | ")
}
