package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeUserHierarchyStructureCmd = &cobra.Command{
	Use:   "describe-user-hierarchy-structure",
	Short: "Describes the hierarchy structure of the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeUserHierarchyStructureCmd).Standalone()

	connect_describeUserHierarchyStructureCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_describeUserHierarchyStructureCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_describeUserHierarchyStructureCmd)
}
