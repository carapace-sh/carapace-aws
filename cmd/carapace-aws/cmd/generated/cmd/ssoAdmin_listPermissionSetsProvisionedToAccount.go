package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listPermissionSetsProvisionedToAccountCmd = &cobra.Command{
	Use:   "list-permission-sets-provisioned-to-account",
	Short: "Lists all the permission sets that are provisioned to a specified Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listPermissionSetsProvisionedToAccountCmd).Standalone()

	ssoAdmin_listPermissionSetsProvisionedToAccountCmd.Flags().String("account-id", "", "The identifier of the Amazon Web Services account from which to list the assignments.")
	ssoAdmin_listPermissionSetsProvisionedToAccountCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_listPermissionSetsProvisionedToAccountCmd.Flags().String("max-results", "", "The maximum number of results to display for the assignment.")
	ssoAdmin_listPermissionSetsProvisionedToAccountCmd.Flags().String("next-token", "", "The pagination token for the list API.")
	ssoAdmin_listPermissionSetsProvisionedToAccountCmd.Flags().String("provisioning-status", "", "The status object for the permission set provisioning operation.")
	ssoAdmin_listPermissionSetsProvisionedToAccountCmd.MarkFlagRequired("account-id")
	ssoAdmin_listPermissionSetsProvisionedToAccountCmd.MarkFlagRequired("instance-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_listPermissionSetsProvisionedToAccountCmd)
}
