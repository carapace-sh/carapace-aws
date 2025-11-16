package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_configureLogsForChannelCmd = &cobra.Command{
	Use:   "configure-logs-for-channel",
	Short: "Configures Amazon CloudWatch log settings for a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_configureLogsForChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_configureLogsForChannelCmd).Standalone()

		mediatailor_configureLogsForChannelCmd.Flags().String("channel-name", "", "The name of the channel.")
		mediatailor_configureLogsForChannelCmd.Flags().String("log-types", "", "The types of logs to collect.")
		mediatailor_configureLogsForChannelCmd.MarkFlagRequired("channel-name")
		mediatailor_configureLogsForChannelCmd.MarkFlagRequired("log-types")
	})
	mediatailorCmd.AddCommand(mediatailor_configureLogsForChannelCmd)
}
