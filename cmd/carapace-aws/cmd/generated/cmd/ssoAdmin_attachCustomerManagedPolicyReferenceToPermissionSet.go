package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_attachCustomerManagedPolicyReferenceToPermissionSetCmd = &cobra.Command{
	Use:   "attach-customer-managed-policy-reference-to-permission-set",
	Short: "Attaches the specified customer managed policy to the specified [PermissionSet]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_attachCustomerManagedPolicyReferenceToPermissionSetCmd).Standalone()

	ssoAdmin_attachCustomerManagedPolicyReferenceToPermissionSetCmd.Flags().String("customer-managed-policy-reference", "", "Specifies the name and path of a customer managed policy.")
	ssoAdmin_attachCustomerManagedPolicyReferenceToPermissionSetCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_attachCustomerManagedPolicyReferenceToPermissionSetCmd.Flags().String("permission-set-arn", "", "The ARN of the `PermissionSet`.")
	ssoAdmin_attachCustomerManagedPolicyReferenceToPermissionSetCmd.MarkFlagRequired("customer-managed-policy-reference")
	ssoAdmin_attachCustomerManagedPolicyReferenceToPermissionSetCmd.MarkFlagRequired("instance-arn")
	ssoAdmin_attachCustomerManagedPolicyReferenceToPermissionSetCmd.MarkFlagRequired("permission-set-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_attachCustomerManagedPolicyReferenceToPermissionSetCmd)
}
