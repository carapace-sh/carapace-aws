package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_describePermissionSetProvisioningStatusCmd = &cobra.Command{
	Use:   "describe-permission-set-provisioning-status",
	Short: "Describes the status for the given permission set provisioning request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_describePermissionSetProvisioningStatusCmd).Standalone()

	ssoAdmin_describePermissionSetProvisioningStatusCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_describePermissionSetProvisioningStatusCmd.Flags().String("provision-permission-set-request-id", "", "The identifier that is provided by the [ProvisionPermissionSet]() call to retrieve the current status of the provisioning workflow.")
	ssoAdmin_describePermissionSetProvisioningStatusCmd.MarkFlagRequired("instance-arn")
	ssoAdmin_describePermissionSetProvisioningStatusCmd.MarkFlagRequired("provision-permission-set-request-id")
	ssoAdminCmd.AddCommand(ssoAdmin_describePermissionSetProvisioningStatusCmd)
}
