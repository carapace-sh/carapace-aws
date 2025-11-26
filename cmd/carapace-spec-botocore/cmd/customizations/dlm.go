package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed dlm/aws.dlm.create-default-role.yaml
var dlm_createDefaultRole []byte

func init() {
	customizations["dlm"] = func(cmd *command.Command) error {
		var specCommand command.Command
		if err := yaml.Unmarshal(dlm_createDefaultRole, &specCommand); err != nil {
			return err
		}
		cmd.Commands = append(cmd.Commands, specCommand)
		return nil
	}
}
