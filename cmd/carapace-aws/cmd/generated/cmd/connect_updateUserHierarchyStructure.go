package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateUserHierarchyStructureCmd = &cobra.Command{
	Use:   "update-user-hierarchy-structure",
	Short: "Updates the user hierarchy structure: add, remove, and rename user hierarchy levels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateUserHierarchyStructureCmd).Standalone()

	connect_updateUserHierarchyStructureCmd.Flags().String("hierarchy-structure", "", "The hierarchy levels to update.")
	connect_updateUserHierarchyStructureCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateUserHierarchyStructureCmd.MarkFlagRequired("hierarchy-structure")
	connect_updateUserHierarchyStructureCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_updateUserHierarchyStructureCmd)
}
