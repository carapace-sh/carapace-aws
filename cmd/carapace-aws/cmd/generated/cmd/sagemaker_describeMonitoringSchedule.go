package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeMonitoringScheduleCmd = &cobra.Command{
	Use:   "describe-monitoring-schedule",
	Short: "Describes the schedule for a monitoring job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeMonitoringScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeMonitoringScheduleCmd).Standalone()

		sagemaker_describeMonitoringScheduleCmd.Flags().String("monitoring-schedule-name", "", "Name of a previously created monitoring schedule.")
		sagemaker_describeMonitoringScheduleCmd.MarkFlagRequired("monitoring-schedule-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeMonitoringScheduleCmd)
}
