package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getSmsTemplateCmd = &cobra.Command{
	Use:   "get-sms-template",
	Short: "Retrieves the content and settings of a message template for messages that are sent through the SMS channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getSmsTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getSmsTemplateCmd).Standalone()

		pinpoint_getSmsTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
		pinpoint_getSmsTemplateCmd.Flags().String("version", "", "The unique identifier for the version of the message template to update, retrieve information about, or delete.")
		pinpoint_getSmsTemplateCmd.MarkFlagRequired("template-name")
	})
	pinpointCmd.AddCommand(pinpoint_getSmsTemplateCmd)
}
