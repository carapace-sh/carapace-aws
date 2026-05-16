package actions

import (
	"github.com/carapace-sh/carapace-aws/pkg/actions/aws"
	spec "github.com/carapace-sh/carapace-spec"
)

// Macros exposes carapace-aws actions for reuse by other commands and specs.
var Macros = spec.MacroMap{
	"Profiles": {
		Name:        "Profiles",
		Description: "Complete AWS profile names from the local AWS config.",
		Function:    "github.com/carapace-sh/carapace-aws/pkg/actions/aws#ActionProfiles",
		Macro:       spec.MacroN(aws.ActionProfiles).Macro,
	},
	"Regions": {
		Name:        "Regions",
		Description: "Complete AWS region names.",
		Function:    "github.com/carapace-sh/carapace-aws/pkg/actions/aws#ActionRegions",
		Macro:       spec.MacroN(aws.ActionRegions).Macro,
	},
}

func init() {
	for name, macro := range Macros {
		spec.AddMacro(name, macro)
	}
}
