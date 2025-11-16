package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_removeNotificationChannelCmd = &cobra.Command{
	Use:   "remove-notification-channel",
	Short: "Remove one anomaly notifications channel for a profiling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_removeNotificationChannelCmd).Standalone()

	codeguruprofiler_removeNotificationChannelCmd.Flags().String("channel-id", "", "The id of the channel that we want to stop receiving notifications.")
	codeguruprofiler_removeNotificationChannelCmd.Flags().String("profiling-group-name", "", "The name of the profiling group we want to change notification configuration for.")
	codeguruprofiler_removeNotificationChannelCmd.MarkFlagRequired("channel-id")
	codeguruprofiler_removeNotificationChannelCmd.MarkFlagRequired("profiling-group-name")
	codeguruprofilerCmd.AddCommand(codeguruprofiler_removeNotificationChannelCmd)
}
