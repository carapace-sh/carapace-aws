package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getEmailTemplateCmd = &cobra.Command{
	Use:   "get-email-template",
	Short: "Retrieves the content and settings of a message template for messages that are sent through the email channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getEmailTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getEmailTemplateCmd).Standalone()

		pinpoint_getEmailTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
		pinpoint_getEmailTemplateCmd.Flags().String("version", "", "The unique identifier for the version of the message template to update, retrieve information about, or delete.")
		pinpoint_getEmailTemplateCmd.MarkFlagRequired("template-name")
	})
	pinpointCmd.AddCommand(pinpoint_getEmailTemplateCmd)
}
