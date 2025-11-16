package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_listPermissionGroupsCmd = &cobra.Command{
	Use:   "list-permission-groups",
	Short: "Lists all available permission groups in FinSpace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_listPermissionGroupsCmd).Standalone()

	finspaceData_listPermissionGroupsCmd.Flags().String("max-results", "", "The maximum number of results per page.")
	finspaceData_listPermissionGroupsCmd.Flags().String("next-token", "", "A token that indicates where a results page should begin.")
	finspaceData_listPermissionGroupsCmd.MarkFlagRequired("max-results")
	finspaceDataCmd.AddCommand(finspaceData_listPermissionGroupsCmd)
}
