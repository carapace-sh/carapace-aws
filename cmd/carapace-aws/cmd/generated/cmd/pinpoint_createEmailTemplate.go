package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_createEmailTemplateCmd = &cobra.Command{
	Use:   "create-email-template",
	Short: "Creates a message template for messages that are sent through the email channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_createEmailTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_createEmailTemplateCmd).Standalone()

		pinpoint_createEmailTemplateCmd.Flags().String("email-template-request", "", "")
		pinpoint_createEmailTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
		pinpoint_createEmailTemplateCmd.MarkFlagRequired("email-template-request")
		pinpoint_createEmailTemplateCmd.MarkFlagRequired("template-name")
	})
	pinpointCmd.AddCommand(pinpoint_createEmailTemplateCmd)
}
