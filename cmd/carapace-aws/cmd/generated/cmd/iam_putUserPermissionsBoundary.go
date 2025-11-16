package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_putUserPermissionsBoundaryCmd = &cobra.Command{
	Use:   "put-user-permissions-boundary",
	Short: "Adds or updates the policy that is specified as the IAM user's permissions boundary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_putUserPermissionsBoundaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_putUserPermissionsBoundaryCmd).Standalone()

		iam_putUserPermissionsBoundaryCmd.Flags().String("permissions-boundary", "", "The ARN of the managed policy that is used to set the permissions boundary for the user.")
		iam_putUserPermissionsBoundaryCmd.Flags().String("user-name", "", "The name (friendly name, not ARN) of the IAM user for which you want to set the permissions boundary.")
		iam_putUserPermissionsBoundaryCmd.MarkFlagRequired("permissions-boundary")
		iam_putUserPermissionsBoundaryCmd.MarkFlagRequired("user-name")
	})
	iamCmd.AddCommand(iam_putUserPermissionsBoundaryCmd)
}
