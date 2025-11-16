package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getSipMediaApplicationLoggingConfigurationCmd = &cobra.Command{
	Use:   "get-sip-media-application-logging-configuration",
	Short: "Retrieves the logging configuration for the specified SIP media application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getSipMediaApplicationLoggingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_getSipMediaApplicationLoggingConfigurationCmd).Standalone()

		chimeSdkVoice_getSipMediaApplicationLoggingConfigurationCmd.Flags().String("sip-media-application-id", "", "The SIP media application ID.")
		chimeSdkVoice_getSipMediaApplicationLoggingConfigurationCmd.MarkFlagRequired("sip-media-application-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getSipMediaApplicationLoggingConfigurationCmd)
}
