package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_putVoiceConnectorExternalSystemsConfigurationCmd = &cobra.Command{
	Use:   "put-voice-connector-external-systems-configuration",
	Short: "Adds an external systems configuration to a Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_putVoiceConnectorExternalSystemsConfigurationCmd).Standalone()

	chimeSdkVoice_putVoiceConnectorExternalSystemsConfigurationCmd.Flags().String("contact-center-system-types", "", "The contact center system to use.")
	chimeSdkVoice_putVoiceConnectorExternalSystemsConfigurationCmd.Flags().String("session-border-controller-types", "", "The session border controllers to use.")
	chimeSdkVoice_putVoiceConnectorExternalSystemsConfigurationCmd.Flags().String("voice-connector-id", "", "The ID of the Voice Connector for which to add the external system configuration.")
	chimeSdkVoice_putVoiceConnectorExternalSystemsConfigurationCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_putVoiceConnectorExternalSystemsConfigurationCmd)
}
