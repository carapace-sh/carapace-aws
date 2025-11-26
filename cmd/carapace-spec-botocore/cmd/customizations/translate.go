package customizations

import "github.com/carapace-sh/carapace-spec/pkg/command"

func init() {
	customizations["translate.translate-document"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "document-content",
			Description: "The content and content type for the document to be translated.",
			Value:       true,
			Required:    true,
		})
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
