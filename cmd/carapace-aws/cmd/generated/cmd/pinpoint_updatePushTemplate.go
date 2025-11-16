package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updatePushTemplateCmd = &cobra.Command{
	Use:   "update-push-template",
	Short: "Updates an existing message template for messages that are sent through a push notification channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updatePushTemplateCmd).Standalone()

	pinpoint_updatePushTemplateCmd.Flags().String("create-new-version", "", "Specifies whether to save the updates as a new version of the message template.")
	pinpoint_updatePushTemplateCmd.Flags().String("push-notification-template-request", "", "")
	pinpoint_updatePushTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
	pinpoint_updatePushTemplateCmd.Flags().String("version", "", "The unique identifier for the version of the message template to update, retrieve information about, or delete.")
	pinpoint_updatePushTemplateCmd.MarkFlagRequired("push-notification-template-request")
	pinpoint_updatePushTemplateCmd.MarkFlagRequired("template-name")
	pinpointCmd.AddCommand(pinpoint_updatePushTemplateCmd)
}
