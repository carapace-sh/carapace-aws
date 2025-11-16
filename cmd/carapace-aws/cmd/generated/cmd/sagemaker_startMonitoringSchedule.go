package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_startMonitoringScheduleCmd = &cobra.Command{
	Use:   "start-monitoring-schedule",
	Short: "Starts a previously stopped monitoring schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_startMonitoringScheduleCmd).Standalone()

	sagemaker_startMonitoringScheduleCmd.Flags().String("monitoring-schedule-name", "", "The name of the schedule to start.")
	sagemaker_startMonitoringScheduleCmd.MarkFlagRequired("monitoring-schedule-name")
	sagemakerCmd.AddCommand(sagemaker_startMonitoringScheduleCmd)
}
