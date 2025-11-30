package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed s3/aws.s3.yaml
var s3 []byte

func init() {
	customizations["s3"] = func(cmd *command.Command) error {
		cmd.Name = "s3api"
		return nil
	}
	customizations["s3.select-object-content"] = func(cmd *command.Command) error {
		delete(cmd.Flags, "--generate-cli-skeleton")
		delete(cmd.Flags, "--cli-input-json=")
		delete(cmd.Flags, "--cli-input-yaml=")
		return nil
	}

	customizations[""] = func(cmd *command.Command) error {
		var specCommand command.Command
		if err := yaml.Unmarshal(s3, &specCommand); err != nil {
			return err
		}
		cmd.Commands = append(cmd.Commands, specCommand)
		return nil
	}
}
