package customizations

import "github.com/carapace-sh/carapace-spec/pkg/command"

func init() {
	customizations["codedeploy"] = func(cmd *command.Command) error {
		cmd.Name = "deploy"
		return nil
	}
}
