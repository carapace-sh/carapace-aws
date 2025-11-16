package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteVoiceTemplateCmd = &cobra.Command{
	Use:   "delete-voice-template",
	Short: "Deletes a message template for messages that were sent through the voice channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteVoiceTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_deleteVoiceTemplateCmd).Standalone()

		pinpoint_deleteVoiceTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
		pinpoint_deleteVoiceTemplateCmd.Flags().String("version", "", "The unique identifier for the version of the message template to update, retrieve information about, or delete.")
		pinpoint_deleteVoiceTemplateCmd.MarkFlagRequired("template-name")
	})
	pinpointCmd.AddCommand(pinpoint_deleteVoiceTemplateCmd)
}
