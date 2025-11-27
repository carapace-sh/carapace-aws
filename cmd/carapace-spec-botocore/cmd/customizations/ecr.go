package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed ecr/aws.ecr.get-login-password.yaml
var ecr_getLoginPassword []byte

func init() {
	customizations["ecr"] = func(cmd *command.Command) error {
		var specCommand command.Command
		if err := yaml.Unmarshal(ecr_getLoginPassword, &specCommand); err != nil {
			return err
		}
		cmd.Commands = append(cmd.Commands, specCommand)
		return nil
	}
}
