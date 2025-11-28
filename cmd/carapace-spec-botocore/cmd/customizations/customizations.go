package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

var customizations = make(map[string]func(cmd *command.Command) error)

func CustomizeCommand(id string, cmd *command.Command) error {
	if f, ok := customizations[id]; ok {
		return f(cmd)
	}
	return nil
}

//go:embed _/aws.yaml
var root []byte

func init() {
	customizations[""] = func(cmd *command.Command) error {
		var specCommand command.Command
		if err := yaml.Unmarshal(root, &specCommand); err != nil {
			return err
		}
		cmd.Description = specCommand.Description
		cmd.Flags = specCommand.Flags
		cmd.PersistentFlags = specCommand.PersistentFlags
		cmd.Completion = specCommand.Completion
		cmd.Documentation = specCommand.Documentation
		return nil

	}
}
