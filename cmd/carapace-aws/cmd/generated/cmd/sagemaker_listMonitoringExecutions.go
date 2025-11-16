package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listMonitoringExecutionsCmd = &cobra.Command{
	Use:   "list-monitoring-executions",
	Short: "Returns list of all monitoring job executions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listMonitoringExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listMonitoringExecutionsCmd).Standalone()

		sagemaker_listMonitoringExecutionsCmd.Flags().String("creation-time-after", "", "A filter that returns only jobs created after a specified time.")
		sagemaker_listMonitoringExecutionsCmd.Flags().String("creation-time-before", "", "A filter that returns only jobs created before a specified time.")
		sagemaker_listMonitoringExecutionsCmd.Flags().String("endpoint-name", "", "Name of a specific endpoint to fetch jobs for.")
		sagemaker_listMonitoringExecutionsCmd.Flags().String("last-modified-time-after", "", "A filter that returns only jobs modified before a specified time.")
		sagemaker_listMonitoringExecutionsCmd.Flags().String("last-modified-time-before", "", "A filter that returns only jobs modified after a specified time.")
		sagemaker_listMonitoringExecutionsCmd.Flags().String("max-results", "", "The maximum number of jobs to return in the response.")
		sagemaker_listMonitoringExecutionsCmd.Flags().String("monitoring-job-definition-name", "", "Gets a list of the monitoring job runs of the specified monitoring job definitions.")
		sagemaker_listMonitoringExecutionsCmd.Flags().String("monitoring-schedule-name", "", "Name of a specific schedule to fetch jobs for.")
		sagemaker_listMonitoringExecutionsCmd.Flags().String("monitoring-type-equals", "", "A filter that returns only the monitoring job runs of the specified monitoring type.")
		sagemaker_listMonitoringExecutionsCmd.Flags().String("next-token", "", "The token returned if the response is truncated.")
		sagemaker_listMonitoringExecutionsCmd.Flags().String("scheduled-time-after", "", "Filter for jobs scheduled after a specified time.")
		sagemaker_listMonitoringExecutionsCmd.Flags().String("scheduled-time-before", "", "Filter for jobs scheduled before a specified time.")
		sagemaker_listMonitoringExecutionsCmd.Flags().String("sort-by", "", "Whether to sort the results by the `Status`, `CreationTime`, or `ScheduledTime` field.")
		sagemaker_listMonitoringExecutionsCmd.Flags().String("sort-order", "", "Whether to sort the results in `Ascending` or `Descending` order.")
		sagemaker_listMonitoringExecutionsCmd.Flags().String("status-equals", "", "A filter that retrieves only jobs with a specific status.")
	})
	sagemakerCmd.AddCommand(sagemaker_listMonitoringExecutionsCmd)
}
