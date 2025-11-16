package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_putSipMediaApplicationLoggingConfigurationCmd = &cobra.Command{
	Use:   "put-sip-media-application-logging-configuration",
	Short: "Updates the logging configuration for the specified SIP media application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_putSipMediaApplicationLoggingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_putSipMediaApplicationLoggingConfigurationCmd).Standalone()

		chimeSdkVoice_putSipMediaApplicationLoggingConfigurationCmd.Flags().String("sip-media-application-id", "", "The SIP media application ID.")
		chimeSdkVoice_putSipMediaApplicationLoggingConfigurationCmd.Flags().String("sip-media-application-logging-configuration", "", "The logging configuration for the specified SIP media application.")
		chimeSdkVoice_putSipMediaApplicationLoggingConfigurationCmd.MarkFlagRequired("sip-media-application-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_putSipMediaApplicationLoggingConfigurationCmd)
}
