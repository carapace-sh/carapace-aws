package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed cloudformation/aws.cloudformation.deploy.yaml
var deployCmd []byte

//go:embed cloudformation/aws.cloudformation.package.yaml
var packageCmd []byte

func init() {
	customizations["cloudformation"] = func(cmd *command.Command) error {
		var deployCommand command.Command
		if err := yaml.Unmarshal(deployCmd, &deployCommand); err != nil {
			return err
		}
		var packageCommand command.Command
		if err := yaml.Unmarshal(packageCmd, &packageCommand); err != nil {
			return err
		}
		cmd.Commands = append(cmd.Commands, deployCommand, packageCommand)
		return nil
	}
}
