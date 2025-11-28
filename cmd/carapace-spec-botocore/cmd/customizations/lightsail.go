package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed lightsail/aws.lightsail.push-container-image.yaml
var lightsail_pushContainerImage []byte

func init() {
	customizations["lightsail"] = func(cmd *command.Command) error {
		var specCommand command.Command
		if err := yaml.Unmarshal(lightsail_pushContainerImage, &specCommand); err != nil {
			return err
		}
		cmd.Commands = append(cmd.Commands, specCommand)
		return nil
	}
}
