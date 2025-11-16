package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_listPermissionGroupsByUserCmd = &cobra.Command{
	Use:   "list-permission-groups-by-user",
	Short: "Lists all the permission groups that are associated with a specific user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_listPermissionGroupsByUserCmd).Standalone()

	finspaceData_listPermissionGroupsByUserCmd.Flags().String("max-results", "", "The maximum number of results per page.")
	finspaceData_listPermissionGroupsByUserCmd.Flags().String("next-token", "", "A token that indicates where a results page should begin.")
	finspaceData_listPermissionGroupsByUserCmd.Flags().String("user-id", "", "The unique identifier for the user.")
	finspaceData_listPermissionGroupsByUserCmd.MarkFlagRequired("max-results")
	finspaceData_listPermissionGroupsByUserCmd.MarkFlagRequired("user-id")
	finspaceDataCmd.AddCommand(finspaceData_listPermissionGroupsByUserCmd)
}
