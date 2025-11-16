package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pi_listAvailableResourceMetricsCmd = &cobra.Command{
	Use:   "list-available-resource-metrics",
	Short: "Retrieve metrics of the specified types that can be queried for a specified DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pi_listAvailableResourceMetricsCmd).Standalone()

	pi_listAvailableResourceMetricsCmd.Flags().String("identifier", "", "An immutable identifier for a data source that is unique within an Amazon Web Services Region.")
	pi_listAvailableResourceMetricsCmd.Flags().String("max-results", "", "The maximum number of items to return.")
	pi_listAvailableResourceMetricsCmd.Flags().String("metric-types", "", "The types of metrics to return in the response.")
	pi_listAvailableResourceMetricsCmd.Flags().String("next-token", "", "An optional pagination token provided by a previous request.")
	pi_listAvailableResourceMetricsCmd.Flags().String("service-type", "", "The Amazon Web Services service for which Performance Insights returns metrics.")
	pi_listAvailableResourceMetricsCmd.MarkFlagRequired("identifier")
	pi_listAvailableResourceMetricsCmd.MarkFlagRequired("metric-types")
	pi_listAvailableResourceMetricsCmd.MarkFlagRequired("service-type")
	piCmd.AddCommand(pi_listAvailableResourceMetricsCmd)
}
