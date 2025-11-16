package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getVoiceConnectorExternalSystemsConfigurationCmd = &cobra.Command{
	Use:   "get-voice-connector-external-systems-configuration",
	Short: "Gets information about an external systems configuration for a Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getVoiceConnectorExternalSystemsConfigurationCmd).Standalone()

	chimeSdkVoice_getVoiceConnectorExternalSystemsConfigurationCmd.Flags().String("voice-connector-id", "", "The ID of the Voice Connector for which to return information about the external system configuration.")
	chimeSdkVoice_getVoiceConnectorExternalSystemsConfigurationCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getVoiceConnectorExternalSystemsConfigurationCmd)
}
