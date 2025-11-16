package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateEmailTemplateCmd = &cobra.Command{
	Use:   "update-email-template",
	Short: "Updates an existing message template for messages that are sent through the email channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateEmailTemplateCmd).Standalone()

	pinpoint_updateEmailTemplateCmd.Flags().String("create-new-version", "", "Specifies whether to save the updates as a new version of the message template.")
	pinpoint_updateEmailTemplateCmd.Flags().String("email-template-request", "", "")
	pinpoint_updateEmailTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
	pinpoint_updateEmailTemplateCmd.Flags().String("version", "", "The unique identifier for the version of the message template to update, retrieve information about, or delete.")
	pinpoint_updateEmailTemplateCmd.MarkFlagRequired("email-template-request")
	pinpoint_updateEmailTemplateCmd.MarkFlagRequired("template-name")
	pinpointCmd.AddCommand(pinpoint_updateEmailTemplateCmd)
}
