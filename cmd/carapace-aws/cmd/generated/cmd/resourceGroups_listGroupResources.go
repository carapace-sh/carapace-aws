package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_listGroupResourcesCmd = &cobra.Command{
	Use:   "list-group-resources",
	Short: "Returns a list of Amazon resource names (ARNs) of the resources that are members of a specified resource group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_listGroupResourcesCmd).Standalone()

	resourceGroups_listGroupResourcesCmd.Flags().String("filters", "", "Filters, formatted as [ResourceFilter]() objects, that you want to apply to a `ListGroupResources` operation.")
	resourceGroups_listGroupResourcesCmd.Flags().String("group", "", "The name or the Amazon resource name (ARN) of the resource group.")
	resourceGroups_listGroupResourcesCmd.Flags().String("group-name", "", "***Deprecated - don't use this parameter.")
	resourceGroups_listGroupResourcesCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
	resourceGroups_listGroupResourcesCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	resourceGroupsCmd.AddCommand(resourceGroups_listGroupResourcesCmd)
}
