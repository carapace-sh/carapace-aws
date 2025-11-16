package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_getPermissionsBoundaryForPermissionSetCmd = &cobra.Command{
	Use:   "get-permissions-boundary-for-permission-set",
	Short: "Obtains the permissions boundary for a specified [PermissionSet]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_getPermissionsBoundaryForPermissionSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_getPermissionsBoundaryForPermissionSetCmd).Standalone()

		ssoAdmin_getPermissionsBoundaryForPermissionSetCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
		ssoAdmin_getPermissionsBoundaryForPermissionSetCmd.Flags().String("permission-set-arn", "", "The ARN of the `PermissionSet`.")
		ssoAdmin_getPermissionsBoundaryForPermissionSetCmd.MarkFlagRequired("instance-arn")
		ssoAdmin_getPermissionsBoundaryForPermissionSetCmd.MarkFlagRequired("permission-set-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_getPermissionsBoundaryForPermissionSetCmd)
}
