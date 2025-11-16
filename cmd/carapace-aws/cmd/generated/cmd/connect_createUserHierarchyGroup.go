package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createUserHierarchyGroupCmd = &cobra.Command{
	Use:   "create-user-hierarchy-group",
	Short: "Creates a new user hierarchy group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createUserHierarchyGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_createUserHierarchyGroupCmd).Standalone()

		connect_createUserHierarchyGroupCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_createUserHierarchyGroupCmd.Flags().String("name", "", "The name of the user hierarchy group.")
		connect_createUserHierarchyGroupCmd.Flags().String("parent-group-id", "", "The identifier for the parent hierarchy group.")
		connect_createUserHierarchyGroupCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		connect_createUserHierarchyGroupCmd.MarkFlagRequired("instance-id")
		connect_createUserHierarchyGroupCmd.MarkFlagRequired("name")
	})
	connectCmd.AddCommand(connect_createUserHierarchyGroupCmd)
}
