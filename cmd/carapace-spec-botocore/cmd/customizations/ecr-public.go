package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed ecr-public/aws.ecr-public.get-login-password.yaml
var ecrPublic_getLoginPassword []byte

func init() {
	customizations["ecr-public"] = func(cmd *command.Command) error {
		var specCommand command.Command
		if err := yaml.Unmarshal(ecrPublic_getLoginPassword, &specCommand); err != nil {
			return err
		}
		cmd.Commands = append(cmd.Commands, specCommand)
		return nil
	}
}
