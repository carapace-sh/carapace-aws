package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed cloudfront/aws.cloudfront.sign.yaml
var signCmd []byte

func init() {
	customizations["cloudfront"] = func(cmd *command.Command) error {
		var signCommand command.Command
		if err := yaml.Unmarshal(signCmd, &signCommand); err != nil {
			return err
		}
		cmd.Commands = append(cmd.Commands, signCommand)
		return nil
	}

	customizations["cloudfront.create-invalidation"] = func(cmd *command.Command) error {
		// TODO needs more customizations
		cmd.AddFlag(command.Flag{
			Longhand:    "paths",
			Description: "The space-separated paths to be invalidated.",
			Value:       true,
		})
		return nil
	}

	customizations["cloudfront.create-distribution"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "origin-domain-name",
			Description: "The domain name for your origin.",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "default-root-object",
			Description: "The object that you want CloudFront to return.",
			Value:       true,
		})
		return nil
	}

	customizations["cloudfront.update-distribution"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "default-root-object",
			Description: "The object that you want CloudFront to return.",
			Value:       true,
		})
		return nil
	}
}
