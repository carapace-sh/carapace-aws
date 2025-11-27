package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed events/aws.events.wizard.yaml
var events_wizard []byte

func init() {
	customizations["events"] = func(cmd *command.Command) error {
		var specCommand command.Command
		if err := yaml.Unmarshal(events_wizard, &specCommand); err != nil {
			return err
		}
		cmd.Commands = append(cmd.Commands, specCommand)
		return nil
	}
}
