package customizations

import (
	"github.com/carapace-sh/carapace-spec/pkg/command"
)

func init() {
	customizations["agenttoolkit"] = func(cmd *command.Command) error {
		cmd.Name = "agent-toolkit"
		return nil
	}
}
