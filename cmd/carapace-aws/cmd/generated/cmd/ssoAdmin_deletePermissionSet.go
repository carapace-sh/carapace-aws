package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_deletePermissionSetCmd = &cobra.Command{
	Use:   "delete-permission-set",
	Short: "Deletes the specified permission set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_deletePermissionSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_deletePermissionSetCmd).Standalone()

		ssoAdmin_deletePermissionSetCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
		ssoAdmin_deletePermissionSetCmd.Flags().String("permission-set-arn", "", "The ARN of the permission set that should be deleted.")
		ssoAdmin_deletePermissionSetCmd.MarkFlagRequired("instance-arn")
		ssoAdmin_deletePermissionSetCmd.MarkFlagRequired("permission-set-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_deletePermissionSetCmd)
}
