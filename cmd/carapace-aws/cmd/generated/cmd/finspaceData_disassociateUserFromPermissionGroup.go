package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_disassociateUserFromPermissionGroupCmd = &cobra.Command{
	Use:   "disassociate-user-from-permission-group",
	Short: "Removes a user from a permission group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_disassociateUserFromPermissionGroupCmd).Standalone()

	finspaceData_disassociateUserFromPermissionGroupCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspaceData_disassociateUserFromPermissionGroupCmd.Flags().String("permission-group-id", "", "The unique identifier for the permission group.")
	finspaceData_disassociateUserFromPermissionGroupCmd.Flags().String("user-id", "", "The unique identifier for the user.")
	finspaceData_disassociateUserFromPermissionGroupCmd.MarkFlagRequired("permission-group-id")
	finspaceData_disassociateUserFromPermissionGroupCmd.MarkFlagRequired("user-id")
	finspaceDataCmd.AddCommand(finspaceData_disassociateUserFromPermissionGroupCmd)
}
