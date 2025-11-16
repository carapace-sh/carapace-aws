package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_updateIdentityProviderCmd = &cobra.Command{
	Use:   "update-identity-provider",
	Short: "Modifies the configuration and trust relationship between a third-party identity provider (IdP) and a user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_updateIdentityProviderCmd).Standalone()

	cognitoIdp_updateIdentityProviderCmd.Flags().String("attribute-mapping", "", "A mapping of IdP attributes to standard and custom user pool attributes.")
	cognitoIdp_updateIdentityProviderCmd.Flags().String("idp-identifiers", "", "An array of IdP identifiers, for example `\"IdPIdentifiers\": [ \"MyIdP\", \"MyIdP2\" ]`.")
	cognitoIdp_updateIdentityProviderCmd.Flags().String("provider-details", "", "The scopes, URLs, and identifiers for your external identity provider.")
	cognitoIdp_updateIdentityProviderCmd.Flags().String("provider-name", "", "The name of the IdP that you want to update.")
	cognitoIdp_updateIdentityProviderCmd.Flags().String("user-pool-id", "", "The Id of the user pool where you want to update your IdP.")
	cognitoIdp_updateIdentityProviderCmd.MarkFlagRequired("provider-name")
	cognitoIdp_updateIdentityProviderCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_updateIdentityProviderCmd)
}
