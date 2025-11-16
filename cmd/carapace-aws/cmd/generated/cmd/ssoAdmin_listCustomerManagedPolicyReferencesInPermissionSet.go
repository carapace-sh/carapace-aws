package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listCustomerManagedPolicyReferencesInPermissionSetCmd = &cobra.Command{
	Use:   "list-customer-managed-policy-references-in-permission-set",
	Short: "Lists all customer managed policies attached to a specified [PermissionSet]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listCustomerManagedPolicyReferencesInPermissionSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_listCustomerManagedPolicyReferencesInPermissionSetCmd).Standalone()

		ssoAdmin_listCustomerManagedPolicyReferencesInPermissionSetCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
		ssoAdmin_listCustomerManagedPolicyReferencesInPermissionSetCmd.Flags().String("max-results", "", "The maximum number of results to display for the list call.")
		ssoAdmin_listCustomerManagedPolicyReferencesInPermissionSetCmd.Flags().String("next-token", "", "The pagination token for the list API.")
		ssoAdmin_listCustomerManagedPolicyReferencesInPermissionSetCmd.Flags().String("permission-set-arn", "", "The ARN of the `PermissionSet`.")
		ssoAdmin_listCustomerManagedPolicyReferencesInPermissionSetCmd.MarkFlagRequired("instance-arn")
		ssoAdmin_listCustomerManagedPolicyReferencesInPermissionSetCmd.MarkFlagRequired("permission-set-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_listCustomerManagedPolicyReferencesInPermissionSetCmd)
}
