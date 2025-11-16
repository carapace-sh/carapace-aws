package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_putVoiceConnectorLoggingConfigurationCmd = &cobra.Command{
	Use:   "put-voice-connector-logging-configuration",
	Short: "Updates a Voice Connector's logging configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_putVoiceConnectorLoggingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_putVoiceConnectorLoggingConfigurationCmd).Standalone()

		chimeSdkVoice_putVoiceConnectorLoggingConfigurationCmd.Flags().String("logging-configuration", "", "The logging configuration being updated.")
		chimeSdkVoice_putVoiceConnectorLoggingConfigurationCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
		chimeSdkVoice_putVoiceConnectorLoggingConfigurationCmd.MarkFlagRequired("logging-configuration")
		chimeSdkVoice_putVoiceConnectorLoggingConfigurationCmd.MarkFlagRequired("voice-connector-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_putVoiceConnectorLoggingConfigurationCmd)
}
