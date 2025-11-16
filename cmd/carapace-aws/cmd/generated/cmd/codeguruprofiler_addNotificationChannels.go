package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_addNotificationChannelsCmd = &cobra.Command{
	Use:   "add-notification-channels",
	Short: "Add up to 2 anomaly notifications channels for a profiling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_addNotificationChannelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruprofiler_addNotificationChannelsCmd).Standalone()

		codeguruprofiler_addNotificationChannelsCmd.Flags().String("channels", "", "One or 2 channels to report to when anomalies are detected.")
		codeguruprofiler_addNotificationChannelsCmd.Flags().String("profiling-group-name", "", "The name of the profiling group that we are setting up notifications for.")
		codeguruprofiler_addNotificationChannelsCmd.MarkFlagRequired("channels")
		codeguruprofiler_addNotificationChannelsCmd.MarkFlagRequired("profiling-group-name")
	})
	codeguruprofilerCmd.AddCommand(codeguruprofiler_addNotificationChannelsCmd)
}
