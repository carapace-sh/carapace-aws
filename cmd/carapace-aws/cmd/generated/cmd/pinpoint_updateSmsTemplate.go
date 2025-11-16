package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateSmsTemplateCmd = &cobra.Command{
	Use:   "update-sms-template",
	Short: "Updates an existing message template for messages that are sent through the SMS channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateSmsTemplateCmd).Standalone()

	pinpoint_updateSmsTemplateCmd.Flags().String("create-new-version", "", "Specifies whether to save the updates as a new version of the message template.")
	pinpoint_updateSmsTemplateCmd.Flags().String("smstemplate-request", "", "")
	pinpoint_updateSmsTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
	pinpoint_updateSmsTemplateCmd.Flags().String("version", "", "The unique identifier for the version of the message template to update, retrieve information about, or delete.")
	pinpoint_updateSmsTemplateCmd.MarkFlagRequired("smstemplate-request")
	pinpoint_updateSmsTemplateCmd.MarkFlagRequired("template-name")
	pinpointCmd.AddCommand(pinpoint_updateSmsTemplateCmd)
}
