package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_deleteVoiceConnectorCmd = &cobra.Command{
	Use:   "delete-voice-connector",
	Short: "Deletes an Amazon Chime SDK Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_deleteVoiceConnectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_deleteVoiceConnectorCmd).Standalone()

		chimeSdkVoice_deleteVoiceConnectorCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
		chimeSdkVoice_deleteVoiceConnectorCmd.MarkFlagRequired("voice-connector-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_deleteVoiceConnectorCmd)
}
