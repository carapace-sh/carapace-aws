package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateUserHierarchyGroupNameCmd = &cobra.Command{
	Use:   "update-user-hierarchy-group-name",
	Short: "Updates the name of the user hierarchy group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateUserHierarchyGroupNameCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateUserHierarchyGroupNameCmd).Standalone()

		connect_updateUserHierarchyGroupNameCmd.Flags().String("hierarchy-group-id", "", "The identifier of the hierarchy group.")
		connect_updateUserHierarchyGroupNameCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateUserHierarchyGroupNameCmd.Flags().String("name", "", "The name of the hierarchy group.")
		connect_updateUserHierarchyGroupNameCmd.MarkFlagRequired("hierarchy-group-id")
		connect_updateUserHierarchyGroupNameCmd.MarkFlagRequired("instance-id")
		connect_updateUserHierarchyGroupNameCmd.MarkFlagRequired("name")
	})
	connectCmd.AddCommand(connect_updateUserHierarchyGroupNameCmd)
}
