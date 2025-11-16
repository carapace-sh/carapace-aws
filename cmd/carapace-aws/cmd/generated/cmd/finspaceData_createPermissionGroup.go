package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_createPermissionGroupCmd = &cobra.Command{
	Use:   "create-permission-group",
	Short: "Creates a group of permissions for various actions that a user can perform in FinSpace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_createPermissionGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceData_createPermissionGroupCmd).Standalone()

		finspaceData_createPermissionGroupCmd.Flags().String("application-permissions", "", "The option to indicate FinSpace application permissions that are granted to a specific group.")
		finspaceData_createPermissionGroupCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
		finspaceData_createPermissionGroupCmd.Flags().String("description", "", "A brief description for the permission group.")
		finspaceData_createPermissionGroupCmd.Flags().String("name", "", "The name of the permission group.")
		finspaceData_createPermissionGroupCmd.MarkFlagRequired("application-permissions")
		finspaceData_createPermissionGroupCmd.MarkFlagRequired("name")
	})
	finspaceDataCmd.AddCommand(finspaceData_createPermissionGroupCmd)
}
