package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_getNotificationConfigurationCmd = &cobra.Command{
	Use:   "get-notification-configuration",
	Short: "Get the current configuration for anomaly notifications for a profiling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_getNotificationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruprofiler_getNotificationConfigurationCmd).Standalone()

		codeguruprofiler_getNotificationConfigurationCmd.Flags().String("profiling-group-name", "", "The name of the profiling group we want to get the notification configuration for.")
		codeguruprofiler_getNotificationConfigurationCmd.MarkFlagRequired("profiling-group-name")
	})
	codeguruprofilerCmd.AddCommand(codeguruprofiler_getNotificationConfigurationCmd)
}
