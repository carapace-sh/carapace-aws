package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getPushTemplateCmd = &cobra.Command{
	Use:   "get-push-template",
	Short: "Retrieves the content and settings of a message template for messages that are sent through a push notification channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getPushTemplateCmd).Standalone()

	pinpoint_getPushTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
	pinpoint_getPushTemplateCmd.Flags().String("version", "", "The unique identifier for the version of the message template to update, retrieve information about, or delete.")
	pinpoint_getPushTemplateCmd.MarkFlagRequired("template-name")
	pinpointCmd.AddCommand(pinpoint_getPushTemplateCmd)
}
