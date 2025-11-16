package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_createVoiceTemplateCmd = &cobra.Command{
	Use:   "create-voice-template",
	Short: "Creates a message template for messages that are sent through the voice channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_createVoiceTemplateCmd).Standalone()

	pinpoint_createVoiceTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
	pinpoint_createVoiceTemplateCmd.Flags().String("voice-template-request", "", "")
	pinpoint_createVoiceTemplateCmd.MarkFlagRequired("template-name")
	pinpoint_createVoiceTemplateCmd.MarkFlagRequired("voice-template-request")
	pinpointCmd.AddCommand(pinpoint_createVoiceTemplateCmd)
}
