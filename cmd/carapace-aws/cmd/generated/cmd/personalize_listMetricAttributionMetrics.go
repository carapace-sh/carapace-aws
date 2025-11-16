package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_listMetricAttributionMetricsCmd = &cobra.Command{
	Use:   "list-metric-attribution-metrics",
	Short: "Lists the metrics for the metric attribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_listMetricAttributionMetricsCmd).Standalone()

	personalize_listMetricAttributionMetricsCmd.Flags().String("max-results", "", "The maximum number of metrics to return in one page of results.")
	personalize_listMetricAttributionMetricsCmd.Flags().String("metric-attribution-arn", "", "The Amazon Resource Name (ARN) of the metric attribution to retrieve attributes for.")
	personalize_listMetricAttributionMetricsCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
	personalizeCmd.AddCommand(personalize_listMetricAttributionMetricsCmd)
}
