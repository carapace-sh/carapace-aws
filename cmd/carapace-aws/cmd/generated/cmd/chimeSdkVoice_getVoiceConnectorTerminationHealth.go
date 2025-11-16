package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getVoiceConnectorTerminationHealthCmd = &cobra.Command{
	Use:   "get-voice-connector-termination-health",
	Short: "Retrieves information about the last time a `SIP OPTIONS` ping was received from your SIP infrastructure for the specified Amazon Chime SDK Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getVoiceConnectorTerminationHealthCmd).Standalone()

	chimeSdkVoice_getVoiceConnectorTerminationHealthCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_getVoiceConnectorTerminationHealthCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getVoiceConnectorTerminationHealthCmd)
}
