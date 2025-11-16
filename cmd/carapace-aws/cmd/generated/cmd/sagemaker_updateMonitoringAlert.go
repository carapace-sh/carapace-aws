package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateMonitoringAlertCmd = &cobra.Command{
	Use:   "update-monitoring-alert",
	Short: "Update the parameters of a model monitor alert.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateMonitoringAlertCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateMonitoringAlertCmd).Standalone()

		sagemaker_updateMonitoringAlertCmd.Flags().String("datapoints-to-alert", "", "Within `EvaluationPeriod`, how many execution failures will raise an alert.")
		sagemaker_updateMonitoringAlertCmd.Flags().String("evaluation-period", "", "The number of most recent monitoring executions to consider when evaluating alert status.")
		sagemaker_updateMonitoringAlertCmd.Flags().String("monitoring-alert-name", "", "The name of a monitoring alert.")
		sagemaker_updateMonitoringAlertCmd.Flags().String("monitoring-schedule-name", "", "The name of a monitoring schedule.")
		sagemaker_updateMonitoringAlertCmd.MarkFlagRequired("datapoints-to-alert")
		sagemaker_updateMonitoringAlertCmd.MarkFlagRequired("evaluation-period")
		sagemaker_updateMonitoringAlertCmd.MarkFlagRequired("monitoring-alert-name")
		sagemaker_updateMonitoringAlertCmd.MarkFlagRequired("monitoring-schedule-name")
	})
	sagemakerCmd.AddCommand(sagemaker_updateMonitoringAlertCmd)
}
