package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listManagedPoliciesInPermissionSetCmd = &cobra.Command{
	Use:   "list-managed-policies-in-permission-set",
	Short: "Lists the Amazon Web Services managed policy that is attached to a specified permission set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listManagedPoliciesInPermissionSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_listManagedPoliciesInPermissionSetCmd).Standalone()

		ssoAdmin_listManagedPoliciesInPermissionSetCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
		ssoAdmin_listManagedPoliciesInPermissionSetCmd.Flags().String("max-results", "", "The maximum number of results to display for the [PermissionSet]().")
		ssoAdmin_listManagedPoliciesInPermissionSetCmd.Flags().String("next-token", "", "The pagination token for the list API.")
		ssoAdmin_listManagedPoliciesInPermissionSetCmd.Flags().String("permission-set-arn", "", "The ARN of the [PermissionSet]() whose managed policies will be listed.")
		ssoAdmin_listManagedPoliciesInPermissionSetCmd.MarkFlagRequired("instance-arn")
		ssoAdmin_listManagedPoliciesInPermissionSetCmd.MarkFlagRequired("permission-set-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_listManagedPoliciesInPermissionSetCmd)
}
