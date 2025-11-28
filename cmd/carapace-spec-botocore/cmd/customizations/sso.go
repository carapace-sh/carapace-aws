package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed sso/aws.sso.login.yaml
var sso_login []byte

func init() {
	customizations["sso"] = func(cmd *command.Command) error {
		var specCommand command.Command
		if err := yaml.Unmarshal(sso_login, &specCommand); err != nil {
			return err
		}
		cmd.Commands = append(cmd.Commands, specCommand)
		return nil
	}

	customizations["sso.logout"] = func(cmd *command.Command) error {
		delete(cmd.Flags, "--access-token=!")
		delete(cmd.Flags, "--cli-input-json=")
		delete(cmd.Flags, "--cli-input-yaml=")
		delete(cmd.Flags, "--generate-cli-skeleton")
		return nil
	}
}
