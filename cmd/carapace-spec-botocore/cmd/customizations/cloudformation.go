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

// TODO handle this in a generic customization function
func AddCloudformationSubcommands(cloudformationCommand *command.Command) error {
	var deployCommand command.Command
	if err := yaml.Unmarshal(deployCmd, &deployCommand); err != nil {
		return err
	}
	var packageCommand command.Command
	if err := yaml.Unmarshal(packageCmd, &packageCommand); err != nil {
		return err
	}
	cloudformationCommand.Commands = append(cloudformationCommand.Commands, deployCommand, packageCommand)
	return nil
}
