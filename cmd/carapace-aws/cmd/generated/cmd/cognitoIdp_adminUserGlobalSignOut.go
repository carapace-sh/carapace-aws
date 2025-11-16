package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminUserGlobalSignOutCmd = &cobra.Command{
	Use:   "admin-user-global-sign-out",
	Short: "Invalidates the identity, access, and refresh tokens that Amazon Cognito issued to a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminUserGlobalSignOutCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_adminUserGlobalSignOutCmd).Standalone()

		cognitoIdp_adminUserGlobalSignOutCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to sign out a user.")
		cognitoIdp_adminUserGlobalSignOutCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
		cognitoIdp_adminUserGlobalSignOutCmd.MarkFlagRequired("user-pool-id")
		cognitoIdp_adminUserGlobalSignOutCmd.MarkFlagRequired("username")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_adminUserGlobalSignOutCmd)
}
