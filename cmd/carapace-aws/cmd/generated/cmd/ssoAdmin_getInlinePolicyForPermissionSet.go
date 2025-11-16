package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_getInlinePolicyForPermissionSetCmd = &cobra.Command{
	Use:   "get-inline-policy-for-permission-set",
	Short: "Obtains the inline policy assigned to the permission set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_getInlinePolicyForPermissionSetCmd).Standalone()

	ssoAdmin_getInlinePolicyForPermissionSetCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_getInlinePolicyForPermissionSetCmd.Flags().String("permission-set-arn", "", "The ARN of the permission set.")
	ssoAdmin_getInlinePolicyForPermissionSetCmd.MarkFlagRequired("instance-arn")
	ssoAdmin_getInlinePolicyForPermissionSetCmd.MarkFlagRequired("permission-set-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_getInlinePolicyForPermissionSetCmd)
}
