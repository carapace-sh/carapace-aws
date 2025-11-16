package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_stopMonitoringScheduleCmd = &cobra.Command{
	Use:   "stop-monitoring-schedule",
	Short: "Stops a previously started monitoring schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_stopMonitoringScheduleCmd).Standalone()

	sagemaker_stopMonitoringScheduleCmd.Flags().String("monitoring-schedule-name", "", "The name of the schedule to stop.")
	sagemaker_stopMonitoringScheduleCmd.MarkFlagRequired("monitoring-schedule-name")
	sagemakerCmd.AddCommand(sagemaker_stopMonitoringScheduleCmd)
}
