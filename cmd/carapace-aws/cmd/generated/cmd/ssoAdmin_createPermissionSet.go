package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_createPermissionSetCmd = &cobra.Command{
	Use:   "create-permission-set",
	Short: "Creates a permission set within a specified IAM Identity Center instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_createPermissionSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_createPermissionSetCmd).Standalone()

		ssoAdmin_createPermissionSetCmd.Flags().String("description", "", "The description of the [PermissionSet]().")
		ssoAdmin_createPermissionSetCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
		ssoAdmin_createPermissionSetCmd.Flags().String("name", "", "The name of the [PermissionSet]().")
		ssoAdmin_createPermissionSetCmd.Flags().String("relay-state", "", "Used to redirect users within the application during the federation authentication process.")
		ssoAdmin_createPermissionSetCmd.Flags().String("session-duration", "", "The length of time that the application user sessions are valid in the ISO-8601 standard.")
		ssoAdmin_createPermissionSetCmd.Flags().String("tags", "", "The tags to attach to the new [PermissionSet]().")
		ssoAdmin_createPermissionSetCmd.MarkFlagRequired("instance-arn")
		ssoAdmin_createPermissionSetCmd.MarkFlagRequired("name")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_createPermissionSetCmd)
}
