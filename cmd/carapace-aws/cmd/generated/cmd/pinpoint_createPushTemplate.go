package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_createPushTemplateCmd = &cobra.Command{
	Use:   "create-push-template",
	Short: "Creates a message template for messages that are sent through a push notification channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_createPushTemplateCmd).Standalone()

	pinpoint_createPushTemplateCmd.Flags().String("push-notification-template-request", "", "")
	pinpoint_createPushTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
	pinpoint_createPushTemplateCmd.MarkFlagRequired("push-notification-template-request")
	pinpoint_createPushTemplateCmd.MarkFlagRequired("template-name")
	pinpointCmd.AddCommand(pinpoint_createPushTemplateCmd)
}
