package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed cloudfront/aws.cloudfront.sign.yaml
var signCmd []byte

// TODO handle this in a generic customization function
func AddCloudfrontSubcommands(cloudfrontCommand *command.Command) error {
	var signCommand command.Command
	if err := yaml.Unmarshal(signCmd, &signCommand); err != nil {
		return err
	}
	cloudfrontCommand.Commands = append(cloudfrontCommand.Commands, signCommand)
	return nil
}
