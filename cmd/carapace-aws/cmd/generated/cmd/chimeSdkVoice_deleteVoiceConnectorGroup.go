package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_deleteVoiceConnectorGroupCmd = &cobra.Command{
	Use:   "delete-voice-connector-group",
	Short: "Deletes an Amazon Chime SDK Voice Connector group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_deleteVoiceConnectorGroupCmd).Standalone()

	chimeSdkVoice_deleteVoiceConnectorGroupCmd.Flags().String("voice-connector-group-id", "", "The Voice Connector Group ID.")
	chimeSdkVoice_deleteVoiceConnectorGroupCmd.MarkFlagRequired("voice-connector-group-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_deleteVoiceConnectorGroupCmd)
}
