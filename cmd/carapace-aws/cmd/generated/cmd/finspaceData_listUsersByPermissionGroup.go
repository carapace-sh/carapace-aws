package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_listUsersByPermissionGroupCmd = &cobra.Command{
	Use:   "list-users-by-permission-group",
	Short: "Lists details of all the users in a specific permission group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_listUsersByPermissionGroupCmd).Standalone()

	finspaceData_listUsersByPermissionGroupCmd.Flags().String("max-results", "", "The maximum number of results per page.")
	finspaceData_listUsersByPermissionGroupCmd.Flags().String("next-token", "", "A token that indicates where a results page should begin.")
	finspaceData_listUsersByPermissionGroupCmd.Flags().String("permission-group-id", "", "The unique identifier for the permission group.")
	finspaceData_listUsersByPermissionGroupCmd.MarkFlagRequired("max-results")
	finspaceData_listUsersByPermissionGroupCmd.MarkFlagRequired("permission-group-id")
	finspaceDataCmd.AddCommand(finspaceData_listUsersByPermissionGroupCmd)
}
