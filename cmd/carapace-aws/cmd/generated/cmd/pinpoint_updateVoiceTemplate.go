package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateVoiceTemplateCmd = &cobra.Command{
	Use:   "update-voice-template",
	Short: "Updates an existing message template for messages that are sent through the voice channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateVoiceTemplateCmd).Standalone()

	pinpoint_updateVoiceTemplateCmd.Flags().String("create-new-version", "", "Specifies whether to save the updates as a new version of the message template.")
	pinpoint_updateVoiceTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
	pinpoint_updateVoiceTemplateCmd.Flags().String("version", "", "The unique identifier for the version of the message template to update, retrieve information about, or delete.")
	pinpoint_updateVoiceTemplateCmd.Flags().String("voice-template-request", "", "")
	pinpoint_updateVoiceTemplateCmd.MarkFlagRequired("template-name")
	pinpoint_updateVoiceTemplateCmd.MarkFlagRequired("voice-template-request")
	pinpointCmd.AddCommand(pinpoint_updateVoiceTemplateCmd)
}
