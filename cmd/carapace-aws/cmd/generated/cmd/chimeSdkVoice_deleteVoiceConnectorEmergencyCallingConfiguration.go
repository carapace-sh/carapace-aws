package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_deleteVoiceConnectorEmergencyCallingConfigurationCmd = &cobra.Command{
	Use:   "delete-voice-connector-emergency-calling-configuration",
	Short: "Deletes the emergency calling details from the specified Amazon Chime SDK Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_deleteVoiceConnectorEmergencyCallingConfigurationCmd).Standalone()

	chimeSdkVoice_deleteVoiceConnectorEmergencyCallingConfigurationCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_deleteVoiceConnectorEmergencyCallingConfigurationCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_deleteVoiceConnectorEmergencyCallingConfigurationCmd)
}
