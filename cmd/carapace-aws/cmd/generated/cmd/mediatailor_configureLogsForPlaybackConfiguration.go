package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_configureLogsForPlaybackConfigurationCmd = &cobra.Command{
	Use:   "configure-logs-for-playback-configuration",
	Short: "Defines where AWS Elemental MediaTailor sends logs for the playback configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_configureLogsForPlaybackConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_configureLogsForPlaybackConfigurationCmd).Standalone()

		mediatailor_configureLogsForPlaybackConfigurationCmd.Flags().String("ads-interaction-log", "", "The event types that MediaTailor emits in logs for interactions with the ADS.")
		mediatailor_configureLogsForPlaybackConfigurationCmd.Flags().String("enabled-logging-strategies", "", "The method used for collecting logs from AWS Elemental MediaTailor.")
		mediatailor_configureLogsForPlaybackConfigurationCmd.Flags().String("manifest-service-interaction-log", "", "The event types that MediaTailor emits in logs for interactions with the origin server.")
		mediatailor_configureLogsForPlaybackConfigurationCmd.Flags().String("percent-enabled", "", "The percentage of session logs that MediaTailor sends to your CloudWatch Logs account.")
		mediatailor_configureLogsForPlaybackConfigurationCmd.Flags().String("playback-configuration-name", "", "The name of the playback configuration.")
		mediatailor_configureLogsForPlaybackConfigurationCmd.MarkFlagRequired("percent-enabled")
		mediatailor_configureLogsForPlaybackConfigurationCmd.MarkFlagRequired("playback-configuration-name")
	})
	mediatailorCmd.AddCommand(mediatailor_configureLogsForPlaybackConfigurationCmd)
}
