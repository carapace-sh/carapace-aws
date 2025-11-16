package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getAggregateDiscoveredResourceCountsCmd = &cobra.Command{
	Use:   "get-aggregate-discovered-resource-counts",
	Short: "Returns the resource counts across accounts and regions that are present in your Config aggregator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getAggregateDiscoveredResourceCountsCmd).Standalone()

	config_getAggregateDiscoveredResourceCountsCmd.Flags().String("configuration-aggregator-name", "", "The name of the configuration aggregator.")
	config_getAggregateDiscoveredResourceCountsCmd.Flags().String("filters", "", "Filters the results based on the `ResourceCountFilters` object.")
	config_getAggregateDiscoveredResourceCountsCmd.Flags().String("group-by-key", "", "The key to group the resource counts.")
	config_getAggregateDiscoveredResourceCountsCmd.Flags().String("limit", "", "The maximum number of [GroupedResourceCount]() objects returned on each page.")
	config_getAggregateDiscoveredResourceCountsCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
	config_getAggregateDiscoveredResourceCountsCmd.MarkFlagRequired("configuration-aggregator-name")
	configCmd.AddCommand(config_getAggregateDiscoveredResourceCountsCmd)
}
