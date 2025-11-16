package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_deleteInlinePolicyFromPermissionSetCmd = &cobra.Command{
	Use:   "delete-inline-policy-from-permission-set",
	Short: "Deletes the inline policy from a specified permission set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_deleteInlinePolicyFromPermissionSetCmd).Standalone()

	ssoAdmin_deleteInlinePolicyFromPermissionSetCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_deleteInlinePolicyFromPermissionSetCmd.Flags().String("permission-set-arn", "", "The ARN of the permission set that will be used to remove access.")
	ssoAdmin_deleteInlinePolicyFromPermissionSetCmd.MarkFlagRequired("instance-arn")
	ssoAdmin_deleteInlinePolicyFromPermissionSetCmd.MarkFlagRequired("permission-set-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_deleteInlinePolicyFromPermissionSetCmd)
}
