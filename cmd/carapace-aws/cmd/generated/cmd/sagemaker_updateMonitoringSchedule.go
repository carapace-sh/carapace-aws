package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateMonitoringScheduleCmd = &cobra.Command{
	Use:   "update-monitoring-schedule",
	Short: "Updates a previously created schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateMonitoringScheduleCmd).Standalone()

	sagemaker_updateMonitoringScheduleCmd.Flags().String("monitoring-schedule-config", "", "The configuration object that specifies the monitoring schedule and defines the monitoring job.")
	sagemaker_updateMonitoringScheduleCmd.Flags().String("monitoring-schedule-name", "", "The name of the monitoring schedule.")
	sagemaker_updateMonitoringScheduleCmd.MarkFlagRequired("monitoring-schedule-config")
	sagemaker_updateMonitoringScheduleCmd.MarkFlagRequired("monitoring-schedule-name")
	sagemakerCmd.AddCommand(sagemaker_updateMonitoringScheduleCmd)
}
