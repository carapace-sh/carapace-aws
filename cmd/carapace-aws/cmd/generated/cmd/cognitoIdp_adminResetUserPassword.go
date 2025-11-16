package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminResetUserPasswordCmd = &cobra.Command{
	Use:   "admin-reset-user-password",
	Short: "Begins the password reset process.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminResetUserPasswordCmd).Standalone()

	cognitoIdp_adminResetUserPasswordCmd.Flags().String("client-metadata", "", "A map of custom key-value pairs that you can provide as input for any custom workflows that this action triggers.")
	cognitoIdp_adminResetUserPasswordCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to reset the user's password.")
	cognitoIdp_adminResetUserPasswordCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
	cognitoIdp_adminResetUserPasswordCmd.MarkFlagRequired("user-pool-id")
	cognitoIdp_adminResetUserPasswordCmd.MarkFlagRequired("username")
	cognitoIdpCmd.AddCommand(cognitoIdp_adminResetUserPasswordCmd)
}
