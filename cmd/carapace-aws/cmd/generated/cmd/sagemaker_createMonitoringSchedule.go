package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createMonitoringScheduleCmd = &cobra.Command{
	Use:   "create-monitoring-schedule",
	Short: "Creates a schedule that regularly starts Amazon SageMaker AI Processing Jobs to monitor the data captured for an Amazon SageMaker AI Endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createMonitoringScheduleCmd).Standalone()

	sagemaker_createMonitoringScheduleCmd.Flags().String("monitoring-schedule-config", "", "The configuration object that specifies the monitoring schedule and defines the monitoring job.")
	sagemaker_createMonitoringScheduleCmd.Flags().String("monitoring-schedule-name", "", "The name of the monitoring schedule.")
	sagemaker_createMonitoringScheduleCmd.Flags().String("tags", "", "(Optional) An array of key-value pairs.")
	sagemaker_createMonitoringScheduleCmd.MarkFlagRequired("monitoring-schedule-config")
	sagemaker_createMonitoringScheduleCmd.MarkFlagRequired("monitoring-schedule-name")
	sagemakerCmd.AddCommand(sagemaker_createMonitoringScheduleCmd)
}
