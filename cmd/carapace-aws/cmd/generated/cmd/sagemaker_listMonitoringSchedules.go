package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listMonitoringSchedulesCmd = &cobra.Command{
	Use:   "list-monitoring-schedules",
	Short: "Returns list of all monitoring schedules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listMonitoringSchedulesCmd).Standalone()

	sagemaker_listMonitoringSchedulesCmd.Flags().String("creation-time-after", "", "A filter that returns only monitoring schedules created after a specified time.")
	sagemaker_listMonitoringSchedulesCmd.Flags().String("creation-time-before", "", "A filter that returns only monitoring schedules created before a specified time.")
	sagemaker_listMonitoringSchedulesCmd.Flags().String("endpoint-name", "", "Name of a specific endpoint to fetch schedules for.")
	sagemaker_listMonitoringSchedulesCmd.Flags().String("last-modified-time-after", "", "A filter that returns only monitoring schedules modified after a specified time.")
	sagemaker_listMonitoringSchedulesCmd.Flags().String("last-modified-time-before", "", "A filter that returns only monitoring schedules modified before a specified time.")
	sagemaker_listMonitoringSchedulesCmd.Flags().String("max-results", "", "The maximum number of jobs to return in the response.")
	sagemaker_listMonitoringSchedulesCmd.Flags().String("monitoring-job-definition-name", "", "Gets a list of the monitoring schedules for the specified monitoring job definition.")
	sagemaker_listMonitoringSchedulesCmd.Flags().String("monitoring-type-equals", "", "A filter that returns only the monitoring schedules for the specified monitoring type.")
	sagemaker_listMonitoringSchedulesCmd.Flags().String("name-contains", "", "Filter for monitoring schedules whose name contains a specified string.")
	sagemaker_listMonitoringSchedulesCmd.Flags().String("next-token", "", "The token returned if the response is truncated.")
	sagemaker_listMonitoringSchedulesCmd.Flags().String("sort-by", "", "Whether to sort the results by the `Status`, `CreationTime`, or `ScheduledTime` field.")
	sagemaker_listMonitoringSchedulesCmd.Flags().String("sort-order", "", "Whether to sort the results in `Ascending` or `Descending` order.")
	sagemaker_listMonitoringSchedulesCmd.Flags().String("status-equals", "", "A filter that returns only monitoring schedules modified before a specified time.")
	sagemakerCmd.AddCommand(sagemaker_listMonitoringSchedulesCmd)
}
