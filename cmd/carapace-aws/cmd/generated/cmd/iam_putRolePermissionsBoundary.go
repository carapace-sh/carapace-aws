package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_putRolePermissionsBoundaryCmd = &cobra.Command{
	Use:   "put-role-permissions-boundary",
	Short: "Adds or updates the policy that is specified as the IAM role's permissions boundary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_putRolePermissionsBoundaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_putRolePermissionsBoundaryCmd).Standalone()

		iam_putRolePermissionsBoundaryCmd.Flags().String("permissions-boundary", "", "The ARN of the managed policy that is used to set the permissions boundary for the role.")
		iam_putRolePermissionsBoundaryCmd.Flags().String("role-name", "", "The name (friendly name, not ARN) of the IAM role for which you want to set the permissions boundary.")
		iam_putRolePermissionsBoundaryCmd.MarkFlagRequired("permissions-boundary")
		iam_putRolePermissionsBoundaryCmd.MarkFlagRequired("role-name")
	})
	iamCmd.AddCommand(iam_putRolePermissionsBoundaryCmd)
}
