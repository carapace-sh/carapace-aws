package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminSetUserPasswordCmd = &cobra.Command{
	Use:   "admin-set-user-password",
	Short: "Sets the specified user's password in a user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminSetUserPasswordCmd).Standalone()

	cognitoIdp_adminSetUserPasswordCmd.Flags().String("password", "", "The new temporary or permanent password that you want to set for the user.")
	cognitoIdp_adminSetUserPasswordCmd.Flags().String("permanent", "", "Set to `true` to set a password that the user can immediately sign in with.")
	cognitoIdp_adminSetUserPasswordCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to set the user's password.")
	cognitoIdp_adminSetUserPasswordCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
	cognitoIdp_adminSetUserPasswordCmd.MarkFlagRequired("password")
	cognitoIdp_adminSetUserPasswordCmd.MarkFlagRequired("user-pool-id")
	cognitoIdp_adminSetUserPasswordCmd.MarkFlagRequired("username")
	cognitoIdpCmd.AddCommand(cognitoIdp_adminSetUserPasswordCmd)
}
