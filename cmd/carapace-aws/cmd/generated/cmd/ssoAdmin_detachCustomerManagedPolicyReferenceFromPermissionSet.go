package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_detachCustomerManagedPolicyReferenceFromPermissionSetCmd = &cobra.Command{
	Use:   "detach-customer-managed-policy-reference-from-permission-set",
	Short: "Detaches the specified customer managed policy from the specified [PermissionSet]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_detachCustomerManagedPolicyReferenceFromPermissionSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_detachCustomerManagedPolicyReferenceFromPermissionSetCmd).Standalone()

		ssoAdmin_detachCustomerManagedPolicyReferenceFromPermissionSetCmd.Flags().String("customer-managed-policy-reference", "", "Specifies the name and path of a customer managed policy.")
		ssoAdmin_detachCustomerManagedPolicyReferenceFromPermissionSetCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
		ssoAdmin_detachCustomerManagedPolicyReferenceFromPermissionSetCmd.Flags().String("permission-set-arn", "", "The ARN of the `PermissionSet`.")
		ssoAdmin_detachCustomerManagedPolicyReferenceFromPermissionSetCmd.MarkFlagRequired("customer-managed-policy-reference")
		ssoAdmin_detachCustomerManagedPolicyReferenceFromPermissionSetCmd.MarkFlagRequired("instance-arn")
		ssoAdmin_detachCustomerManagedPolicyReferenceFromPermissionSetCmd.MarkFlagRequired("permission-set-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_detachCustomerManagedPolicyReferenceFromPermissionSetCmd)
}
