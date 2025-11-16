package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_deletePermissionsBoundaryFromPermissionSetCmd = &cobra.Command{
	Use:   "delete-permissions-boundary-from-permission-set",
	Short: "Deletes the permissions boundary from a specified [PermissionSet]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_deletePermissionsBoundaryFromPermissionSetCmd).Standalone()

	ssoAdmin_deletePermissionsBoundaryFromPermissionSetCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_deletePermissionsBoundaryFromPermissionSetCmd.Flags().String("permission-set-arn", "", "The ARN of the `PermissionSet`.")
	ssoAdmin_deletePermissionsBoundaryFromPermissionSetCmd.MarkFlagRequired("instance-arn")
	ssoAdmin_deletePermissionsBoundaryFromPermissionSetCmd.MarkFlagRequired("permission-set-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_deletePermissionsBoundaryFromPermissionSetCmd)
}
