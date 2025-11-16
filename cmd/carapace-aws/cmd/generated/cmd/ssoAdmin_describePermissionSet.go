package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_describePermissionSetCmd = &cobra.Command{
	Use:   "describe-permission-set",
	Short: "Gets the details of the permission set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_describePermissionSetCmd).Standalone()

	ssoAdmin_describePermissionSetCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_describePermissionSetCmd.Flags().String("permission-set-arn", "", "The ARN of the permission set.")
	ssoAdmin_describePermissionSetCmd.MarkFlagRequired("instance-arn")
	ssoAdmin_describePermissionSetCmd.MarkFlagRequired("permission-set-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_describePermissionSetCmd)
}
