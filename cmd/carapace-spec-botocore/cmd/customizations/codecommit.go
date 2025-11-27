package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed codecommit/aws.codecommit.credential-helper.yaml
var codecommit_credentialHelper []byte

func init() {
	customizations["codecommit"] = func(cmd *command.Command) error {
		// TODO credential-helper has a `get` subcommand which is currently not tested in the `aws` github action
		var specCommand command.Command
		if err := yaml.Unmarshal(codecommit_credentialHelper, &specCommand); err != nil {
			return err
		}
		cmd.Commands = append(cmd.Commands, specCommand)
		return nil
	}
}
