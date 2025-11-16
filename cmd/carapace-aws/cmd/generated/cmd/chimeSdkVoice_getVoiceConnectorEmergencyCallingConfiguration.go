package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getVoiceConnectorEmergencyCallingConfigurationCmd = &cobra.Command{
	Use:   "get-voice-connector-emergency-calling-configuration",
	Short: "Retrieves the emergency calling configuration details for the specified Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getVoiceConnectorEmergencyCallingConfigurationCmd).Standalone()

	chimeSdkVoice_getVoiceConnectorEmergencyCallingConfigurationCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_getVoiceConnectorEmergencyCallingConfigurationCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getVoiceConnectorEmergencyCallingConfigurationCmd)
}
