package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed codeartifact/aws.codeartifact.login.yaml
var codeartifact_login []byte

func init() {
	customizations["codeartifact"] = func(cmd *command.Command) error {
		var specCommand command.Command
		if err := yaml.Unmarshal(codeartifact_login, &specCommand); err != nil {
			return err
		}
		cmd.Commands = append(cmd.Commands, specCommand)
		return nil
	}
}
