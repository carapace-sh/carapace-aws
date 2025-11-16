package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteMonitoringScheduleCmd = &cobra.Command{
	Use:   "delete-monitoring-schedule",
	Short: "Deletes a monitoring schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteMonitoringScheduleCmd).Standalone()

	sagemaker_deleteMonitoringScheduleCmd.Flags().String("monitoring-schedule-name", "", "The name of the monitoring schedule to delete.")
	sagemaker_deleteMonitoringScheduleCmd.MarkFlagRequired("monitoring-schedule-name")
	sagemakerCmd.AddCommand(sagemaker_deleteMonitoringScheduleCmd)
}
