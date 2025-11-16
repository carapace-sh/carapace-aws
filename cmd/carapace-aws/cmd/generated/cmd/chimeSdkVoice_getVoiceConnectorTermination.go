package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getVoiceConnectorTerminationCmd = &cobra.Command{
	Use:   "get-voice-connector-termination",
	Short: "Retrieves the termination setting details for the specified Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getVoiceConnectorTerminationCmd).Standalone()

	chimeSdkVoice_getVoiceConnectorTerminationCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_getVoiceConnectorTerminationCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getVoiceConnectorTerminationCmd)
}
