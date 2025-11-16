package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteSmsTemplateCmd = &cobra.Command{
	Use:   "delete-sms-template",
	Short: "Deletes a message template for messages that were sent through the SMS channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteSmsTemplateCmd).Standalone()

	pinpoint_deleteSmsTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
	pinpoint_deleteSmsTemplateCmd.Flags().String("version", "", "The unique identifier for the version of the message template to update, retrieve information about, or delete.")
	pinpoint_deleteSmsTemplateCmd.MarkFlagRequired("template-name")
	pinpointCmd.AddCommand(pinpoint_deleteSmsTemplateCmd)
}
