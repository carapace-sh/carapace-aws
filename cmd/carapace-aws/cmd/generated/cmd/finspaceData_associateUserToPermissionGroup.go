package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_associateUserToPermissionGroupCmd = &cobra.Command{
	Use:   "associate-user-to-permission-group",
	Short: "Adds a user to a permission group to grant permissions for actions a user can perform in FinSpace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_associateUserToPermissionGroupCmd).Standalone()

	finspaceData_associateUserToPermissionGroupCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspaceData_associateUserToPermissionGroupCmd.Flags().String("permission-group-id", "", "The unique identifier for the permission group.")
	finspaceData_associateUserToPermissionGroupCmd.Flags().String("user-id", "", "The unique identifier for the user.")
	finspaceData_associateUserToPermissionGroupCmd.MarkFlagRequired("permission-group-id")
	finspaceData_associateUserToPermissionGroupCmd.MarkFlagRequired("user-id")
	finspaceDataCmd.AddCommand(finspaceData_associateUserToPermissionGroupCmd)
}
