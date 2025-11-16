package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_updateVoiceConnectorGroupCmd = &cobra.Command{
	Use:   "update-voice-connector-group",
	Short: "Updates the settings for the specified Amazon Chime SDK Voice Connector group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_updateVoiceConnectorGroupCmd).Standalone()

	chimeSdkVoice_updateVoiceConnectorGroupCmd.Flags().String("name", "", "The name of the Voice Connector group.")
	chimeSdkVoice_updateVoiceConnectorGroupCmd.Flags().String("voice-connector-group-id", "", "The Voice Connector ID.")
	chimeSdkVoice_updateVoiceConnectorGroupCmd.Flags().String("voice-connector-items", "", "The `VoiceConnectorItems` to associate with the Voice Connector group.")
	chimeSdkVoice_updateVoiceConnectorGroupCmd.MarkFlagRequired("name")
	chimeSdkVoice_updateVoiceConnectorGroupCmd.MarkFlagRequired("voice-connector-group-id")
	chimeSdkVoice_updateVoiceConnectorGroupCmd.MarkFlagRequired("voice-connector-items")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_updateVoiceConnectorGroupCmd)
}
