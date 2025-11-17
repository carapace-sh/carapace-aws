package actions

import (
	"github.com/carapace-sh/carapace-aws/pkg/actions/aws"
	spec "github.com/carapace-sh/carapace-spec"
)

func init() {
	// TODO create with go:generate
	spec.AddMacro("Profiles", spec.MacroN(aws.ActionProfiles))
	spec.AddMacro("Regions", spec.MacroN(aws.ActionRegions))
}
