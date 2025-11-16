package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pi_listAvailableResourceDimensionsCmd = &cobra.Command{
	Use:   "list-available-resource-dimensions",
	Short: "Retrieve the dimensions that can be queried for each specified metric type on a specified DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pi_listAvailableResourceDimensionsCmd).Standalone()

	pi_listAvailableResourceDimensionsCmd.Flags().String("authorized-actions", "", "The actions to discover the dimensions you are authorized to access.")
	pi_listAvailableResourceDimensionsCmd.Flags().String("identifier", "", "An immutable identifier for a data source that is unique within an Amazon Web Services Region.")
	pi_listAvailableResourceDimensionsCmd.Flags().String("max-results", "", "The maximum number of items to return in the response.")
	pi_listAvailableResourceDimensionsCmd.Flags().String("metrics", "", "The types of metrics for which to retrieve dimensions.")
	pi_listAvailableResourceDimensionsCmd.Flags().String("next-token", "", "An optional pagination token provided by a previous request.")
	pi_listAvailableResourceDimensionsCmd.Flags().String("service-type", "", "The Amazon Web Services service for which Performance Insights returns metrics.")
	pi_listAvailableResourceDimensionsCmd.MarkFlagRequired("identifier")
	pi_listAvailableResourceDimensionsCmd.MarkFlagRequired("metrics")
	pi_listAvailableResourceDimensionsCmd.MarkFlagRequired("service-type")
	piCmd.AddCommand(pi_listAvailableResourceDimensionsCmd)
}
