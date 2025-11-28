package customizations

import "github.com/carapace-sh/carapace-spec/pkg/command"

func init() {
	customizations["ses.send-email"] = func(cmd *command.Command) error {
		delete(cmd.Flags, "--source=!") // TODO provide remove function in carapace-spec

		cmd.AddFlag(command.Flag{
			Longhand:    "bcc",
			Description: "The email addresses of blind-carbon-copy recipients (Bcc)",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "to",
			Description: "The email addresses of the primary recipients",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "subject",
			Description: "The subject of the message",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "html",
			Description: "The HTML body of the message",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "text",
			Description: "The raw text body of the message",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "from",
			Description: "The email address that is sending the email",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "cc",
			Description: "The email addresses of copy recipients",
			Value:       true,
		})
		return nil
	}
}
