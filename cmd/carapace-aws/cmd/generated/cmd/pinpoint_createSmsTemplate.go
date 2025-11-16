package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_createSmsTemplateCmd = &cobra.Command{
	Use:   "create-sms-template",
	Short: "Creates a message template for messages that are sent through the SMS channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_createSmsTemplateCmd).Standalone()

	pinpoint_createSmsTemplateCmd.Flags().String("smstemplate-request", "", "")
	pinpoint_createSmsTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
	pinpoint_createSmsTemplateCmd.MarkFlagRequired("smstemplate-request")
	pinpoint_createSmsTemplateCmd.MarkFlagRequired("template-name")
	pinpointCmd.AddCommand(pinpoint_createSmsTemplateCmd)
}
