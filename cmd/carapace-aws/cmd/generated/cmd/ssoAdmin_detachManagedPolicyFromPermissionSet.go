package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_detachManagedPolicyFromPermissionSetCmd = &cobra.Command{
	Use:   "detach-managed-policy-from-permission-set",
	Short: "Detaches the attached Amazon Web Services managed policy ARN from the specified permission set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_detachManagedPolicyFromPermissionSetCmd).Standalone()

	ssoAdmin_detachManagedPolicyFromPermissionSetCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_detachManagedPolicyFromPermissionSetCmd.Flags().String("managed-policy-arn", "", "The Amazon Web Services managed policy ARN to be detached from a permission set.")
	ssoAdmin_detachManagedPolicyFromPermissionSetCmd.Flags().String("permission-set-arn", "", "The ARN of the [PermissionSet]() from which the policy should be detached.")
	ssoAdmin_detachManagedPolicyFromPermissionSetCmd.MarkFlagRequired("instance-arn")
	ssoAdmin_detachManagedPolicyFromPermissionSetCmd.MarkFlagRequired("managed-policy-arn")
	ssoAdmin_detachManagedPolicyFromPermissionSetCmd.MarkFlagRequired("permission-set-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_detachManagedPolicyFromPermissionSetCmd)
}
