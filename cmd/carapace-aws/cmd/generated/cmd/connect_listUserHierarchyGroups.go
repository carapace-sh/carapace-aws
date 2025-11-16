package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listUserHierarchyGroupsCmd = &cobra.Command{
	Use:   "list-user-hierarchy-groups",
	Short: "Provides summary information about the hierarchy groups for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listUserHierarchyGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listUserHierarchyGroupsCmd).Standalone()

		connect_listUserHierarchyGroupsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listUserHierarchyGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listUserHierarchyGroupsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listUserHierarchyGroupsCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listUserHierarchyGroupsCmd)
}
