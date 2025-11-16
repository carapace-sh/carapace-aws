package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_deleteVoiceConnectorTerminationCmd = &cobra.Command{
	Use:   "delete-voice-connector-termination",
	Short: "Deletes the termination settings for the specified Amazon Chime SDK Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_deleteVoiceConnectorTerminationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_deleteVoiceConnectorTerminationCmd).Standalone()

		chimeSdkVoice_deleteVoiceConnectorTerminationCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
		chimeSdkVoice_deleteVoiceConnectorTerminationCmd.MarkFlagRequired("voice-connector-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_deleteVoiceConnectorTerminationCmd)
}
