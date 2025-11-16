package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteUserHierarchyGroupCmd = &cobra.Command{
	Use:   "delete-user-hierarchy-group",
	Short: "Deletes an existing user hierarchy group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteUserHierarchyGroupCmd).Standalone()

	connect_deleteUserHierarchyGroupCmd.Flags().String("hierarchy-group-id", "", "The identifier of the hierarchy group.")
	connect_deleteUserHierarchyGroupCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_deleteUserHierarchyGroupCmd.MarkFlagRequired("hierarchy-group-id")
	connect_deleteUserHierarchyGroupCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_deleteUserHierarchyGroupCmd)
}
