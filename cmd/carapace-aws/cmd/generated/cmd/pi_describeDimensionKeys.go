package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pi_describeDimensionKeysCmd = &cobra.Command{
	Use:   "describe-dimension-keys",
	Short: "For a specific time period, retrieve the top `N` dimension keys for a metric.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pi_describeDimensionKeysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pi_describeDimensionKeysCmd).Standalone()

		pi_describeDimensionKeysCmd.Flags().String("additional-metrics", "", "Additional metrics for the top `N` dimension keys.")
		pi_describeDimensionKeysCmd.Flags().String("end-time", "", "The date and time specifying the end of the requested time series data.")
		pi_describeDimensionKeysCmd.Flags().String("filter", "", "One or more filters to apply in the request.")
		pi_describeDimensionKeysCmd.Flags().String("group-by", "", "A specification for how to aggregate the data points from a query result.")
		pi_describeDimensionKeysCmd.Flags().String("identifier", "", "An immutable, Amazon Web Services Region-unique identifier for a data source.")
		pi_describeDimensionKeysCmd.Flags().String("max-results", "", "The maximum number of items to return in the response.")
		pi_describeDimensionKeysCmd.Flags().String("metric", "", "The name of a Performance Insights metric to be measured.")
		pi_describeDimensionKeysCmd.Flags().String("next-token", "", "An optional pagination token provided by a previous request.")
		pi_describeDimensionKeysCmd.Flags().String("partition-by", "", "For each dimension specified in `GroupBy`, specify a secondary dimension to further subdivide the partition keys in the response.")
		pi_describeDimensionKeysCmd.Flags().String("period-in-seconds", "", "The granularity, in seconds, of the data points returned from Performance Insights.")
		pi_describeDimensionKeysCmd.Flags().String("service-type", "", "The Amazon Web Services service for which Performance Insights will return metrics.")
		pi_describeDimensionKeysCmd.Flags().String("start-time", "", "The date and time specifying the beginning of the requested time series data.")
		pi_describeDimensionKeysCmd.MarkFlagRequired("end-time")
		pi_describeDimensionKeysCmd.MarkFlagRequired("group-by")
		pi_describeDimensionKeysCmd.MarkFlagRequired("identifier")
		pi_describeDimensionKeysCmd.MarkFlagRequired("metric")
		pi_describeDimensionKeysCmd.MarkFlagRequired("service-type")
		pi_describeDimensionKeysCmd.MarkFlagRequired("start-time")
	})
	piCmd.AddCommand(pi_describeDimensionKeysCmd)
}
