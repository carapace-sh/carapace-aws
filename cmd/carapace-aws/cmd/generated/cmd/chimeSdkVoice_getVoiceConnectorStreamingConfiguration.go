package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getVoiceConnectorStreamingConfigurationCmd = &cobra.Command{
	Use:   "get-voice-connector-streaming-configuration",
	Short: "Retrieves the streaming configuration details for the specified Amazon Chime SDK Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getVoiceConnectorStreamingConfigurationCmd).Standalone()

	chimeSdkVoice_getVoiceConnectorStreamingConfigurationCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_getVoiceConnectorStreamingConfigurationCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getVoiceConnectorStreamingConfigurationCmd)
}
