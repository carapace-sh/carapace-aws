package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listAccountsForProvisionedPermissionSetCmd = &cobra.Command{
	Use:   "list-accounts-for-provisioned-permission-set",
	Short: "Lists all the Amazon Web Services accounts where the specified permission set is provisioned.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listAccountsForProvisionedPermissionSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_listAccountsForProvisionedPermissionSetCmd).Standalone()

		ssoAdmin_listAccountsForProvisionedPermissionSetCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
		ssoAdmin_listAccountsForProvisionedPermissionSetCmd.Flags().String("max-results", "", "The maximum number of results to display for the [PermissionSet]().")
		ssoAdmin_listAccountsForProvisionedPermissionSetCmd.Flags().String("next-token", "", "The pagination token for the list API.")
		ssoAdmin_listAccountsForProvisionedPermissionSetCmd.Flags().String("permission-set-arn", "", "The ARN of the [PermissionSet]() from which the associated Amazon Web Services accounts will be listed.")
		ssoAdmin_listAccountsForProvisionedPermissionSetCmd.Flags().String("provisioning-status", "", "The permission set provisioning status for an Amazon Web Services account.")
		ssoAdmin_listAccountsForProvisionedPermissionSetCmd.MarkFlagRequired("instance-arn")
		ssoAdmin_listAccountsForProvisionedPermissionSetCmd.MarkFlagRequired("permission-set-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_listAccountsForProvisionedPermissionSetCmd)
}
