package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listMonitoringAlertHistoryCmd = &cobra.Command{
	Use:   "list-monitoring-alert-history",
	Short: "Gets a list of past alerts in a model monitoring schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listMonitoringAlertHistoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listMonitoringAlertHistoryCmd).Standalone()

		sagemaker_listMonitoringAlertHistoryCmd.Flags().String("creation-time-after", "", "A filter that returns only alerts created on or after the specified time.")
		sagemaker_listMonitoringAlertHistoryCmd.Flags().String("creation-time-before", "", "A filter that returns only alerts created on or before the specified time.")
		sagemaker_listMonitoringAlertHistoryCmd.Flags().String("max-results", "", "The maximum number of results to display.")
		sagemaker_listMonitoringAlertHistoryCmd.Flags().String("monitoring-alert-name", "", "The name of a monitoring alert.")
		sagemaker_listMonitoringAlertHistoryCmd.Flags().String("monitoring-schedule-name", "", "The name of a monitoring schedule.")
		sagemaker_listMonitoringAlertHistoryCmd.Flags().String("next-token", "", "If the result of the previous `ListMonitoringAlertHistory` request was truncated, the response includes a `NextToken`.")
		sagemaker_listMonitoringAlertHistoryCmd.Flags().String("sort-by", "", "The field used to sort results.")
		sagemaker_listMonitoringAlertHistoryCmd.Flags().String("sort-order", "", "The sort order, whether `Ascending` or `Descending`, of the alert history.")
		sagemaker_listMonitoringAlertHistoryCmd.Flags().String("status-equals", "", "A filter that retrieves only alerts with a specific status.")
	})
	sagemakerCmd.AddCommand(sagemaker_listMonitoringAlertHistoryCmd)
}
