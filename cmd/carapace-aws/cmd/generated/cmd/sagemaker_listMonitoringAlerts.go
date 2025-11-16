package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listMonitoringAlertsCmd = &cobra.Command{
	Use:   "list-monitoring-alerts",
	Short: "Gets the alerts for a single monitoring schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listMonitoringAlertsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listMonitoringAlertsCmd).Standalone()

		sagemaker_listMonitoringAlertsCmd.Flags().String("max-results", "", "The maximum number of results to display.")
		sagemaker_listMonitoringAlertsCmd.Flags().String("monitoring-schedule-name", "", "The name of a monitoring schedule.")
		sagemaker_listMonitoringAlertsCmd.Flags().String("next-token", "", "If the result of the previous `ListMonitoringAlerts` request was truncated, the response includes a `NextToken`.")
		sagemaker_listMonitoringAlertsCmd.MarkFlagRequired("monitoring-schedule-name")
	})
	sagemakerCmd.AddCommand(sagemaker_listMonitoringAlertsCmd)
}
