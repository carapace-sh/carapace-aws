package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeUserHierarchyGroupCmd = &cobra.Command{
	Use:   "describe-user-hierarchy-group",
	Short: "Describes the specified hierarchy group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeUserHierarchyGroupCmd).Standalone()

	connect_describeUserHierarchyGroupCmd.Flags().String("hierarchy-group-id", "", "The identifier of the hierarchy group.")
	connect_describeUserHierarchyGroupCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_describeUserHierarchyGroupCmd.MarkFlagRequired("hierarchy-group-id")
	connect_describeUserHierarchyGroupCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_describeUserHierarchyGroupCmd)
}
