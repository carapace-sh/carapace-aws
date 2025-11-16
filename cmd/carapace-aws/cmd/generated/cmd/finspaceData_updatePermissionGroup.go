package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_updatePermissionGroupCmd = &cobra.Command{
	Use:   "update-permission-group",
	Short: "Modifies the details of a permission group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_updatePermissionGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceData_updatePermissionGroupCmd).Standalone()

		finspaceData_updatePermissionGroupCmd.Flags().String("application-permissions", "", "The permissions that are granted to a specific group for accessing the FinSpace application.")
		finspaceData_updatePermissionGroupCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
		finspaceData_updatePermissionGroupCmd.Flags().String("description", "", "A brief description for the permission group.")
		finspaceData_updatePermissionGroupCmd.Flags().String("name", "", "The name of the permission group.")
		finspaceData_updatePermissionGroupCmd.Flags().String("permission-group-id", "", "The unique identifier for the permission group to update.")
		finspaceData_updatePermissionGroupCmd.MarkFlagRequired("permission-group-id")
	})
	finspaceDataCmd.AddCommand(finspaceData_updatePermissionGroupCmd)
}
