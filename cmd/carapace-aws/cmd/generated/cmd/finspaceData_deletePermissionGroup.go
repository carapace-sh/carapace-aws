package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_deletePermissionGroupCmd = &cobra.Command{
	Use:   "delete-permission-group",
	Short: "Deletes a permission group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_deletePermissionGroupCmd).Standalone()

	finspaceData_deletePermissionGroupCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspaceData_deletePermissionGroupCmd.Flags().String("permission-group-id", "", "The unique identifier for the permission group that you want to delete.")
	finspaceData_deletePermissionGroupCmd.MarkFlagRequired("permission-group-id")
	finspaceDataCmd.AddCommand(finspaceData_deletePermissionGroupCmd)
}
