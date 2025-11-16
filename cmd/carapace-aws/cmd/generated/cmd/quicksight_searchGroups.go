package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_searchGroupsCmd = &cobra.Command{
	Use:   "search-groups",
	Short: "Use the `SearchGroups` operation to search groups in a specified Quick Sight namespace using the supplied filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_searchGroupsCmd).Standalone()

	quicksight_searchGroupsCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that the group is in.")
	quicksight_searchGroupsCmd.Flags().String("filters", "", "The structure for the search filters that you want to apply to your search.")
	quicksight_searchGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return from this request.")
	quicksight_searchGroupsCmd.Flags().String("namespace", "", "The namespace that you want to search.")
	quicksight_searchGroupsCmd.Flags().String("next-token", "", "A pagination token that can be used in a subsequent request.")
	quicksight_searchGroupsCmd.MarkFlagRequired("aws-account-id")
	quicksight_searchGroupsCmd.MarkFlagRequired("filters")
	quicksight_searchGroupsCmd.MarkFlagRequired("namespace")
	quicksightCmd.AddCommand(quicksight_searchGroupsCmd)
}
