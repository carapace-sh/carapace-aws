package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_deleteVoiceConnectorExternalSystemsConfigurationCmd = &cobra.Command{
	Use:   "delete-voice-connector-external-systems-configuration",
	Short: "Deletes the external systems configuration for a Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_deleteVoiceConnectorExternalSystemsConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_deleteVoiceConnectorExternalSystemsConfigurationCmd).Standalone()

		chimeSdkVoice_deleteVoiceConnectorExternalSystemsConfigurationCmd.Flags().String("voice-connector-id", "", "The ID of the Voice Connector for which to delete the external system configuration.")
		chimeSdkVoice_deleteVoiceConnectorExternalSystemsConfigurationCmd.MarkFlagRequired("voice-connector-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_deleteVoiceConnectorExternalSystemsConfigurationCmd)
}
