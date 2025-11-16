package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listPermissionSetProvisioningStatusCmd = &cobra.Command{
	Use:   "list-permission-set-provisioning-status",
	Short: "Lists the status of the permission set provisioning requests for a specified IAM Identity Center instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listPermissionSetProvisioningStatusCmd).Standalone()

	ssoAdmin_listPermissionSetProvisioningStatusCmd.Flags().String("filter", "", "Filters results based on the passed attribute value.")
	ssoAdmin_listPermissionSetProvisioningStatusCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_listPermissionSetProvisioningStatusCmd.Flags().String("max-results", "", "The maximum number of results to display for the assignment.")
	ssoAdmin_listPermissionSetProvisioningStatusCmd.Flags().String("next-token", "", "The pagination token for the list API.")
	ssoAdmin_listPermissionSetProvisioningStatusCmd.MarkFlagRequired("instance-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_listPermissionSetProvisioningStatusCmd)
}
