package customizations

import "github.com/carapace-sh/carapace-spec/pkg/command"

func init() {
	customizations["config"] = func(cmd *command.Command) error {
		cmd.Name = "configservice"
		return nil
	}
}
