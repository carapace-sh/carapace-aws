package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_getPermissionGroupCmd = &cobra.Command{
	Use:   "get-permission-group",
	Short: "Retrieves the details of a specific permission group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_getPermissionGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceData_getPermissionGroupCmd).Standalone()

		finspaceData_getPermissionGroupCmd.Flags().String("permission-group-id", "", "The unique identifier for the permission group.")
		finspaceData_getPermissionGroupCmd.MarkFlagRequired("permission-group-id")
	})
	finspaceDataCmd.AddCommand(finspaceData_getPermissionGroupCmd)
}
