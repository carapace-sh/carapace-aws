package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_putPermissionsBoundaryToPermissionSetCmd = &cobra.Command{
	Use:   "put-permissions-boundary-to-permission-set",
	Short: "Attaches an Amazon Web Services managed or customer managed policy to the specified [PermissionSet]() as a permissions boundary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_putPermissionsBoundaryToPermissionSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_putPermissionsBoundaryToPermissionSetCmd).Standalone()

		ssoAdmin_putPermissionsBoundaryToPermissionSetCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
		ssoAdmin_putPermissionsBoundaryToPermissionSetCmd.Flags().String("permission-set-arn", "", "The ARN of the `PermissionSet`.")
		ssoAdmin_putPermissionsBoundaryToPermissionSetCmd.Flags().String("permissions-boundary", "", "The permissions boundary that you want to attach to a `PermissionSet`.")
		ssoAdmin_putPermissionsBoundaryToPermissionSetCmd.MarkFlagRequired("instance-arn")
		ssoAdmin_putPermissionsBoundaryToPermissionSetCmd.MarkFlagRequired("permission-set-arn")
		ssoAdmin_putPermissionsBoundaryToPermissionSetCmd.MarkFlagRequired("permissions-boundary")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_putPermissionsBoundaryToPermissionSetCmd)
}
