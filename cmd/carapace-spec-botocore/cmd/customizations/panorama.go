package customizations

import (
	"github.com/carapace-sh/carapace-spec/pkg/command"
)

func init() {
	customizations["panorama"] = func(cmd *command.Command) error {
		return &SkipError{}
	}
}
