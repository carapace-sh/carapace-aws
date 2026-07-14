package customizations

import (
	"github.com/carapace-sh/carapace-spec/pkg/command"
)

func init() {
	customizations["simspaceweaver"] = func(cmd *command.Command) error {
		return &SkipError{}
	}
}
