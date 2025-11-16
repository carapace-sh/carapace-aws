package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getDiscoveredResourceCountsCmd = &cobra.Command{
	Use:   "get-discovered-resource-counts",
	Short: "Returns the resource types, the number of each resource type, and the total number of resources that Config is recording in this region for your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getDiscoveredResourceCountsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_getDiscoveredResourceCountsCmd).Standalone()

		config_getDiscoveredResourceCountsCmd.Flags().String("limit", "", "The maximum number of [ResourceCount]() objects returned on each page.")
		config_getDiscoveredResourceCountsCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
		config_getDiscoveredResourceCountsCmd.Flags().String("resource-types", "", "The comma-separated list that specifies the resource types that you want Config to return (for example, `\"AWS::EC2::Instance\"`, `\"AWS::IAM::User\"`).")
	})
	configCmd.AddCommand(config_getDiscoveredResourceCountsCmd)
}
