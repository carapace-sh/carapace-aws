package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_getIdentityProviderByIdentifierCmd = &cobra.Command{
	Use:   "get-identity-provider-by-identifier",
	Short: "Given the identifier of an identity provider (IdP), for example `examplecorp`, returns information about the user pool configuration for that IdP.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_getIdentityProviderByIdentifierCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_getIdentityProviderByIdentifierCmd).Standalone()

		cognitoIdp_getIdentityProviderByIdentifierCmd.Flags().String("idp-identifier", "", "The identifier that you assigned to your user pool.")
		cognitoIdp_getIdentityProviderByIdentifierCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to get information about the IdP.")
		cognitoIdp_getIdentityProviderByIdentifierCmd.MarkFlagRequired("idp-identifier")
		cognitoIdp_getIdentityProviderByIdentifierCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_getIdentityProviderByIdentifierCmd)
}
