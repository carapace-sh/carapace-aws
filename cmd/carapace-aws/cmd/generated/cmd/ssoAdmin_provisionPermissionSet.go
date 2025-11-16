package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_provisionPermissionSetCmd = &cobra.Command{
	Use:   "provision-permission-set",
	Short: "The process by which a specified permission set is provisioned to the specified target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_provisionPermissionSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_provisionPermissionSetCmd).Standalone()

		ssoAdmin_provisionPermissionSetCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
		ssoAdmin_provisionPermissionSetCmd.Flags().String("permission-set-arn", "", "The ARN of the permission set.")
		ssoAdmin_provisionPermissionSetCmd.Flags().String("target-id", "", "TargetID is an Amazon Web Services account identifier, (For example, 123456789012).")
		ssoAdmin_provisionPermissionSetCmd.Flags().String("target-type", "", "The entity type for which the assignment will be created.")
		ssoAdmin_provisionPermissionSetCmd.MarkFlagRequired("instance-arn")
		ssoAdmin_provisionPermissionSetCmd.MarkFlagRequired("permission-set-arn")
		ssoAdmin_provisionPermissionSetCmd.MarkFlagRequired("target-type")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_provisionPermissionSetCmd)
}
