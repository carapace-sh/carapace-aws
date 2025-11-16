package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_putVoiceConnectorStreamingConfigurationCmd = &cobra.Command{
	Use:   "put-voice-connector-streaming-configuration",
	Short: "Updates a Voice Connector's streaming configuration settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_putVoiceConnectorStreamingConfigurationCmd).Standalone()

	chimeSdkVoice_putVoiceConnectorStreamingConfigurationCmd.Flags().String("streaming-configuration", "", "The streaming settings being updated.")
	chimeSdkVoice_putVoiceConnectorStreamingConfigurationCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_putVoiceConnectorStreamingConfigurationCmd.MarkFlagRequired("streaming-configuration")
	chimeSdkVoice_putVoiceConnectorStreamingConfigurationCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_putVoiceConnectorStreamingConfigurationCmd)
}
