package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed iam/aws.iam.wizard.yaml
var iam_wizard []byte

func init() {
	customizations["iam"] = func(cmd *command.Command) error {
		var specCommand command.Command
		if err := yaml.Unmarshal(iam_wizard, &specCommand); err != nil {
			return err
		}
		cmd.Commands = append(cmd.Commands, specCommand)
		return nil
	}

	customizations["iam.create-virtual-mfa-device"] = func(cmd *command.Command) error {
		if cmd.Completion.Flag == nil {
			cmd.Completion.Flag = map[string][]string{}
		}

		delete(cmd.Flags, "--generate-cli-skeleton")
		delete(cmd.Flags, "--cli-input-json=")
		delete(cmd.Flags, "--cli-input-yaml=")

		cmd.AddFlag(command.Flag{
			Longhand:    "bootstrap-method",
			Description: "Method to use to seed the virtual MFA",
			Value:       true,
			Required:    true,
		})
		cmd.Completion.Flag["bootstrap-method"] = []string{"QRCodePNG", "Base32StringSeed"}
		cmd.Completion.Positional = [][]string{{"$files"}}

		return nil
	}
}
