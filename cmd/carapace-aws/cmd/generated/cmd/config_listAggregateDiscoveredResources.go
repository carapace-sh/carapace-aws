package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_listAggregateDiscoveredResourcesCmd = &cobra.Command{
	Use:   "list-aggregate-discovered-resources",
	Short: "Accepts a resource type and returns a list of resource identifiers that are aggregated for a specific resource type across accounts and regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_listAggregateDiscoveredResourcesCmd).Standalone()

	config_listAggregateDiscoveredResourcesCmd.Flags().String("configuration-aggregator-name", "", "The name of the configuration aggregator.")
	config_listAggregateDiscoveredResourcesCmd.Flags().String("filters", "", "Filters the results based on the `ResourceFilters` object.")
	config_listAggregateDiscoveredResourcesCmd.Flags().String("limit", "", "The maximum number of resource identifiers returned on each page.")
	config_listAggregateDiscoveredResourcesCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
	config_listAggregateDiscoveredResourcesCmd.Flags().String("resource-type", "", "The type of resources that you want Config to list in the response.")
	config_listAggregateDiscoveredResourcesCmd.MarkFlagRequired("configuration-aggregator-name")
	config_listAggregateDiscoveredResourcesCmd.MarkFlagRequired("resource-type")
	configCmd.AddCommand(config_listAggregateDiscoveredResourcesCmd)
}
