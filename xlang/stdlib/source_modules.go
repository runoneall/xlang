package stdlib

import _ "embed"

//go:embed srcmod_enum.x
var enum string

// SourceModules are source type standard library modules.
var SourceModules = map[string]string{
	"enum": enum,
}
