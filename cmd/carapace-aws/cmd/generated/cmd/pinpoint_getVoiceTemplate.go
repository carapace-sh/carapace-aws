package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getVoiceTemplateCmd = &cobra.Command{
	Use:   "get-voice-template",
	Short: "Retrieves the content and settings of a message template for messages that are sent through the voice channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getVoiceTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getVoiceTemplateCmd).Standalone()

		pinpoint_getVoiceTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
		pinpoint_getVoiceTemplateCmd.Flags().String("version", "", "The unique identifier for the version of the message template to update, retrieve information about, or delete.")
		pinpoint_getVoiceTemplateCmd.MarkFlagRequired("template-name")
	})
	pinpointCmd.AddCommand(pinpoint_getVoiceTemplateCmd)
}
