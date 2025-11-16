package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_listGroupingStatusesCmd = &cobra.Command{
	Use:   "list-grouping-statuses",
	Short: "Returns the status of the last grouping or ungrouping action for each resource in the specified application group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_listGroupingStatusesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceGroups_listGroupingStatusesCmd).Standalone()

		resourceGroups_listGroupingStatusesCmd.Flags().String("filters", "", "The filter name and value pair that is used to return more specific results from a list of resources.")
		resourceGroups_listGroupingStatusesCmd.Flags().String("group", "", "The application group identifier, expressed as an Amazon resource name (ARN) or the application group name.")
		resourceGroups_listGroupingStatusesCmd.Flags().String("max-results", "", "The maximum number of resources and their statuses returned in the response.")
		resourceGroups_listGroupingStatusesCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
		resourceGroups_listGroupingStatusesCmd.MarkFlagRequired("group")
	})
	resourceGroupsCmd.AddCommand(resourceGroups_listGroupingStatusesCmd)
}
