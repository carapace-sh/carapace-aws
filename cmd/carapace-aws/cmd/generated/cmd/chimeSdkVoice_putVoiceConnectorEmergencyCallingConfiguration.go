package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_putVoiceConnectorEmergencyCallingConfigurationCmd = &cobra.Command{
	Use:   "put-voice-connector-emergency-calling-configuration",
	Short: "Updates a Voice Connector's emergency calling configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_putVoiceConnectorEmergencyCallingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_putVoiceConnectorEmergencyCallingConfigurationCmd).Standalone()

		chimeSdkVoice_putVoiceConnectorEmergencyCallingConfigurationCmd.Flags().String("emergency-calling-configuration", "", "The configuration being updated.")
		chimeSdkVoice_putVoiceConnectorEmergencyCallingConfigurationCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
		chimeSdkVoice_putVoiceConnectorEmergencyCallingConfigurationCmd.MarkFlagRequired("emergency-calling-configuration")
		chimeSdkVoice_putVoiceConnectorEmergencyCallingConfigurationCmd.MarkFlagRequired("voice-connector-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_putVoiceConnectorEmergencyCallingConfigurationCmd)
}
