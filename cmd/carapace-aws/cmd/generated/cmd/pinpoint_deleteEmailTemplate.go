package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteEmailTemplateCmd = &cobra.Command{
	Use:   "delete-email-template",
	Short: "Deletes a message template for messages that were sent through the email channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteEmailTemplateCmd).Standalone()

	pinpoint_deleteEmailTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
	pinpoint_deleteEmailTemplateCmd.Flags().String("version", "", "The unique identifier for the version of the message template to update, retrieve information about, or delete.")
	pinpoint_deleteEmailTemplateCmd.MarkFlagRequired("template-name")
	pinpointCmd.AddCommand(pinpoint_deleteEmailTemplateCmd)
}
