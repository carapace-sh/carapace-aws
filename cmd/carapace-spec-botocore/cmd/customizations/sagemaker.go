package customizations

import (
	"github.com/carapace-sh/carapace-spec/pkg/command"
)

func init() {
	customizations["sagemaker"] = func(cmd *command.Command) error {
		for i := range cmd.Commands {
			if cmd.Commands[i].Name != "wait" {
				continue
			}
			for j := range cmd.Commands[i].Commands {
				waitSubCmd := &cmd.Commands[i].Commands[j]
				if waitSubCmd.Name != "image-version-created" && waitSubCmd.Name != "image-version-deleted" {
					continue
				}
				delete(waitSubCmd.Flags, "--version-number=")
			}
		}
		return nil
	}
}
