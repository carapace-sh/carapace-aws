package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteUserPermissionsBoundaryCmd = &cobra.Command{
	Use:   "delete-user-permissions-boundary",
	Short: "Deletes the permissions boundary for the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteUserPermissionsBoundaryCmd).Standalone()

	iam_deleteUserPermissionsBoundaryCmd.Flags().String("user-name", "", "The name (friendly name, not ARN) of the IAM user from which you want to remove the permissions boundary.")
	iam_deleteUserPermissionsBoundaryCmd.MarkFlagRequired("user-name")
	iamCmd.AddCommand(iam_deleteUserPermissionsBoundaryCmd)
}
