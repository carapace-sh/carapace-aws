package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_batchGetRumMetricDefinitionsCmd = &cobra.Command{
	Use:   "batch-get-rum-metric-definitions",
	Short: "Retrieves the list of metrics and dimensions that a RUM app monitor is sending to a single destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_batchGetRumMetricDefinitionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rum_batchGetRumMetricDefinitionsCmd).Standalone()

		rum_batchGetRumMetricDefinitionsCmd.Flags().String("app-monitor-name", "", "The name of the CloudWatch RUM app monitor that is sending the metrics.")
		rum_batchGetRumMetricDefinitionsCmd.Flags().String("destination", "", "The type of destination that you want to view metrics for.")
		rum_batchGetRumMetricDefinitionsCmd.Flags().String("destination-arn", "", "This parameter is required if `Destination` is `Evidently`.")
		rum_batchGetRumMetricDefinitionsCmd.Flags().String("max-results", "", "The maximum number of results to return in one operation.")
		rum_batchGetRumMetricDefinitionsCmd.Flags().String("next-token", "", "Use the token returned by the previous operation to request the next page of results.")
		rum_batchGetRumMetricDefinitionsCmd.MarkFlagRequired("app-monitor-name")
		rum_batchGetRumMetricDefinitionsCmd.MarkFlagRequired("destination")
	})
	rumCmd.AddCommand(rum_batchGetRumMetricDefinitionsCmd)
}
