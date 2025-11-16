package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_deleteIdentityProviderCmd = &cobra.Command{
	Use:   "delete-identity-provider",
	Short: "Deletes a user pool identity provider (IdP).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_deleteIdentityProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_deleteIdentityProviderCmd).Standalone()

		cognitoIdp_deleteIdentityProviderCmd.Flags().String("provider-name", "", "The name of the IdP that you want to delete.")
		cognitoIdp_deleteIdentityProviderCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to delete the identity provider.")
		cognitoIdp_deleteIdentityProviderCmd.MarkFlagRequired("provider-name")
		cognitoIdp_deleteIdentityProviderCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_deleteIdentityProviderCmd)
}
