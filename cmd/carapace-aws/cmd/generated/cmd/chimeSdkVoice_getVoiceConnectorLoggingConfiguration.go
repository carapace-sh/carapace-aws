package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getVoiceConnectorLoggingConfigurationCmd = &cobra.Command{
	Use:   "get-voice-connector-logging-configuration",
	Short: "Retrieves the logging configuration settings for the specified Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getVoiceConnectorLoggingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_getVoiceConnectorLoggingConfigurationCmd).Standalone()

		chimeSdkVoice_getVoiceConnectorLoggingConfigurationCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
		chimeSdkVoice_getVoiceConnectorLoggingConfigurationCmd.MarkFlagRequired("voice-connector-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getVoiceConnectorLoggingConfigurationCmd)
}
