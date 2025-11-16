package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminConfirmSignUpCmd = &cobra.Command{
	Use:   "admin-confirm-sign-up",
	Short: "Confirms user sign-up as an administrator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminConfirmSignUpCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_adminConfirmSignUpCmd).Standalone()

		cognitoIdp_adminConfirmSignUpCmd.Flags().String("client-metadata", "", "A map of custom key-value pairs that you can provide as input for any custom workflows that this action triggers.")
		cognitoIdp_adminConfirmSignUpCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to confirm a user's sign-up request.")
		cognitoIdp_adminConfirmSignUpCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
		cognitoIdp_adminConfirmSignUpCmd.MarkFlagRequired("user-pool-id")
		cognitoIdp_adminConfirmSignUpCmd.MarkFlagRequired("username")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_adminConfirmSignUpCmd)
}
