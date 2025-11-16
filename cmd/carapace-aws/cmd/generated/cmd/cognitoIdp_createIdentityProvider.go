package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_createIdentityProviderCmd = &cobra.Command{
	Use:   "create-identity-provider",
	Short: "Adds a configuration and trust relationship between a third-party identity provider (IdP) and a user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_createIdentityProviderCmd).Standalone()

	cognitoIdp_createIdentityProviderCmd.Flags().String("attribute-mapping", "", "A mapping of IdP attributes to standard and custom user pool attributes.")
	cognitoIdp_createIdentityProviderCmd.Flags().String("idp-identifiers", "", "An array of IdP identifiers, for example `\"IdPIdentifiers\": [ \"MyIdP\", \"MyIdP2\" ]`.")
	cognitoIdp_createIdentityProviderCmd.Flags().String("provider-details", "", "The scopes, URLs, and identifiers for your external identity provider.")
	cognitoIdp_createIdentityProviderCmd.Flags().String("provider-name", "", "The name that you want to assign to the IdP.")
	cognitoIdp_createIdentityProviderCmd.Flags().String("provider-type", "", "The type of IdP that you want to add.")
	cognitoIdp_createIdentityProviderCmd.Flags().String("user-pool-id", "", "The Id of the user pool where you want to create an IdP.")
	cognitoIdp_createIdentityProviderCmd.MarkFlagRequired("provider-details")
	cognitoIdp_createIdentityProviderCmd.MarkFlagRequired("provider-name")
	cognitoIdp_createIdentityProviderCmd.MarkFlagRequired("provider-type")
	cognitoIdp_createIdentityProviderCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_createIdentityProviderCmd)
}
