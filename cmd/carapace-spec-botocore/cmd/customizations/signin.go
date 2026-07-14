package customizations

import "github.com/carapace-sh/carapace-spec/pkg/command"

func init() {
	customizations["signin.create-oauth2-token"] = func(cmd *command.Command) error {
		cmd.Aliases = append(cmd.Aliases, "create-o-auth2-token")
		return nil
	}

	customizations["translate.import-terminology"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "data-file",
			Description: "The path to the file of the code you are uploading.",
			Value:       true,
			Required:    true,
		})
		return nil
	}
}
