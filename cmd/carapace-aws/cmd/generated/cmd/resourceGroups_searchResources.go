package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_searchResourcesCmd = &cobra.Command{
	Use:   "search-resources",
	Short: "Returns a list of Amazon Web Services resource identifiers that matches the specified query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_searchResourcesCmd).Standalone()

	resourceGroups_searchResourcesCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
	resourceGroups_searchResourcesCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	resourceGroups_searchResourcesCmd.Flags().String("resource-query", "", "The search query, using the same formats that are supported for resource group definition.")
	resourceGroups_searchResourcesCmd.MarkFlagRequired("resource-query")
	resourceGroupsCmd.AddCommand(resourceGroups_searchResourcesCmd)
}
