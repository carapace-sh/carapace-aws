package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deletePushTemplateCmd = &cobra.Command{
	Use:   "delete-push-template",
	Short: "Deletes a message template for messages that were sent through a push notification channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deletePushTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_deletePushTemplateCmd).Standalone()

		pinpoint_deletePushTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
		pinpoint_deletePushTemplateCmd.Flags().String("version", "", "The unique identifier for the version of the message template to update, retrieve information about, or delete.")
		pinpoint_deletePushTemplateCmd.MarkFlagRequired("template-name")
	})
	pinpointCmd.AddCommand(pinpoint_deletePushTemplateCmd)
}
