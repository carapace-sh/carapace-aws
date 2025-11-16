package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchUserHierarchyGroupsCmd = &cobra.Command{
	Use:   "search-user-hierarchy-groups",
	Short: "Searches UserHierarchyGroups in an Amazon Connect instance, with optional filtering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchUserHierarchyGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_searchUserHierarchyGroupsCmd).Standalone()

		connect_searchUserHierarchyGroupsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_searchUserHierarchyGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_searchUserHierarchyGroupsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_searchUserHierarchyGroupsCmd.Flags().String("search-criteria", "", "The search criteria to be used to return UserHierarchyGroups.")
		connect_searchUserHierarchyGroupsCmd.Flags().String("search-filter", "", "Filters to be applied to search results.")
		connect_searchUserHierarchyGroupsCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_searchUserHierarchyGroupsCmd)
}
