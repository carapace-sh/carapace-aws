package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_putInlinePolicyToPermissionSetCmd = &cobra.Command{
	Use:   "put-inline-policy-to-permission-set",
	Short: "Attaches an inline policy to a permission set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_putInlinePolicyToPermissionSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_putInlinePolicyToPermissionSetCmd).Standalone()

		ssoAdmin_putInlinePolicyToPermissionSetCmd.Flags().String("inline-policy", "", "The inline policy to attach to a [PermissionSet]().")
		ssoAdmin_putInlinePolicyToPermissionSetCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
		ssoAdmin_putInlinePolicyToPermissionSetCmd.Flags().String("permission-set-arn", "", "The ARN of the permission set.")
		ssoAdmin_putInlinePolicyToPermissionSetCmd.MarkFlagRequired("inline-policy")
		ssoAdmin_putInlinePolicyToPermissionSetCmd.MarkFlagRequired("instance-arn")
		ssoAdmin_putInlinePolicyToPermissionSetCmd.MarkFlagRequired("permission-set-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_putInlinePolicyToPermissionSetCmd)
}
