package customizations

import "github.com/carapace-sh/carapace-spec/pkg/command"

var customizations = make(map[string]func(cmd *command.Command) error)

func CustomizeCommand(id string, cmd *command.Command) error {
	if f, ok := customizations[id]; ok {
		return f(cmd)
	}
	return nil
}
