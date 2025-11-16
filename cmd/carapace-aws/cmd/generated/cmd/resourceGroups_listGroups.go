package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_listGroupsCmd = &cobra.Command{
	Use:   "list-groups",
	Short: "Returns a list of existing Resource Groups in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_listGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceGroups_listGroupsCmd).Standalone()

		resourceGroups_listGroupsCmd.Flags().String("filters", "", "Filters, formatted as [GroupFilter]() objects, that you want to apply to a `ListGroups` operation.")
		resourceGroups_listGroupsCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
		resourceGroups_listGroupsCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	})
	resourceGroupsCmd.AddCommand(resourceGroups_listGroupsCmd)
}
