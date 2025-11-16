package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_attachManagedPolicyToPermissionSetCmd = &cobra.Command{
	Use:   "attach-managed-policy-to-permission-set",
	Short: "Attaches an Amazon Web Services managed policy ARN to a permission set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_attachManagedPolicyToPermissionSetCmd).Standalone()

	ssoAdmin_attachManagedPolicyToPermissionSetCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_attachManagedPolicyToPermissionSetCmd.Flags().String("managed-policy-arn", "", "The Amazon Web Services managed policy ARN to be attached to a permission set.")
	ssoAdmin_attachManagedPolicyToPermissionSetCmd.Flags().String("permission-set-arn", "", "The ARN of the [PermissionSet]() that the managed policy should be attached to.")
	ssoAdmin_attachManagedPolicyToPermissionSetCmd.MarkFlagRequired("instance-arn")
	ssoAdmin_attachManagedPolicyToPermissionSetCmd.MarkFlagRequired("managed-policy-arn")
	ssoAdmin_attachManagedPolicyToPermissionSetCmd.MarkFlagRequired("permission-set-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_attachManagedPolicyToPermissionSetCmd)
}
