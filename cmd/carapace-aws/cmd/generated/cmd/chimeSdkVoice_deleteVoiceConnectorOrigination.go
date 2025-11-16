package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_deleteVoiceConnectorOriginationCmd = &cobra.Command{
	Use:   "delete-voice-connector-origination",
	Short: "Deletes the origination settings for the specified Amazon Chime SDK Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_deleteVoiceConnectorOriginationCmd).Standalone()

	chimeSdkVoice_deleteVoiceConnectorOriginationCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_deleteVoiceConnectorOriginationCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_deleteVoiceConnectorOriginationCmd)
}
