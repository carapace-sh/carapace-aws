package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateUserHierarchyCmd = &cobra.Command{
	Use:   "update-user-hierarchy",
	Short: "Assigns the specified hierarchy group to the specified user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateUserHierarchyCmd).Standalone()

	connect_updateUserHierarchyCmd.Flags().String("hierarchy-group-id", "", "The identifier of the hierarchy group.")
	connect_updateUserHierarchyCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateUserHierarchyCmd.Flags().String("user-id", "", "The identifier of the user account.")
	connect_updateUserHierarchyCmd.MarkFlagRequired("instance-id")
	connect_updateUserHierarchyCmd.MarkFlagRequired("user-id")
	connectCmd.AddCommand(connect_updateUserHierarchyCmd)
}
