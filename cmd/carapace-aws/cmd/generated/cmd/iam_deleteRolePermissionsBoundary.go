package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteRolePermissionsBoundaryCmd = &cobra.Command{
	Use:   "delete-role-permissions-boundary",
	Short: "Deletes the permissions boundary for the specified IAM role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteRolePermissionsBoundaryCmd).Standalone()

	iam_deleteRolePermissionsBoundaryCmd.Flags().String("role-name", "", "The name (friendly name, not ARN) of the IAM role from which you want to remove the permissions boundary.")
	iam_deleteRolePermissionsBoundaryCmd.MarkFlagRequired("role-name")
	iamCmd.AddCommand(iam_deleteRolePermissionsBoundaryCmd)
}
