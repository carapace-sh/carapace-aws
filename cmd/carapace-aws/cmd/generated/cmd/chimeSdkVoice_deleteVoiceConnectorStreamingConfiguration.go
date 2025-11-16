package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_deleteVoiceConnectorStreamingConfigurationCmd = &cobra.Command{
	Use:   "delete-voice-connector-streaming-configuration",
	Short: "Deletes a Voice Connector's streaming configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_deleteVoiceConnectorStreamingConfigurationCmd).Standalone()

	chimeSdkVoice_deleteVoiceConnectorStreamingConfigurationCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_deleteVoiceConnectorStreamingConfigurationCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_deleteVoiceConnectorStreamingConfigurationCmd)
}
