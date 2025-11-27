package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed lambda/aws.lambda.wizard.yaml
var lambda_wizard []byte

func init() {
	customizations["lambda"] = func(cmd *command.Command) error {
		var specCommand command.Command
		if err := yaml.Unmarshal(lambda_wizard, &specCommand); err != nil {
			return err
		}
		cmd.Commands = append(cmd.Commands, specCommand)
		return nil
	}

	customizations["lambda.publish-layer-version"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "zip-file",
			Description: "The path to the zip file of the content you are uploading.",
			Value:       true,
		})
		return nil
	}
	customizations["lambda.create-function"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "zip-file",
			Description: "The path to the zip file of the code you are uploading.",
			Value:       true,
		})
		return nil
	}
}
