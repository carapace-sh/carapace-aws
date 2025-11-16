package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_updatePermissionSetCmd = &cobra.Command{
	Use:   "update-permission-set",
	Short: "Updates an existing permission set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_updatePermissionSetCmd).Standalone()

	ssoAdmin_updatePermissionSetCmd.Flags().String("description", "", "The description of the [PermissionSet]().")
	ssoAdmin_updatePermissionSetCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_updatePermissionSetCmd.Flags().String("permission-set-arn", "", "The ARN of the permission set.")
	ssoAdmin_updatePermissionSetCmd.Flags().String("relay-state", "", "Used to redirect users within the application during the federation authentication process.")
	ssoAdmin_updatePermissionSetCmd.Flags().String("session-duration", "", "The length of time that the application user sessions are valid for in the ISO-8601 standard.")
	ssoAdmin_updatePermissionSetCmd.MarkFlagRequired("instance-arn")
	ssoAdmin_updatePermissionSetCmd.MarkFlagRequired("permission-set-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_updatePermissionSetCmd)
}
