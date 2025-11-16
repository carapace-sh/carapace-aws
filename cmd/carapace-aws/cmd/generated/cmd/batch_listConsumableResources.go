package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_listConsumableResourcesCmd = &cobra.Command{
	Use:   "list-consumable-resources",
	Short: "Returns a list of Batch consumable resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_listConsumableResourcesCmd).Standalone()

	batch_listConsumableResourcesCmd.Flags().String("filters", "", "The filters to apply to the consumable resource list query.")
	batch_listConsumableResourcesCmd.Flags().String("max-results", "", "The maximum number of results returned by `ListConsumableResources` in paginated output.")
	batch_listConsumableResourcesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListConsumableResources` request where `maxResults` was used and the results exceeded the value of that parameter.")
	batchCmd.AddCommand(batch_listConsumableResourcesCmd)
}
