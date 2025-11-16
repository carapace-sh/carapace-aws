package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_describeIdentityProviderCmd = &cobra.Command{
	Use:   "describe-identity-provider",
	Short: "Given a user pool ID and identity provider (IdP) name, returns details about the IdP.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_describeIdentityProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_describeIdentityProviderCmd).Standalone()

		cognitoIdp_describeIdentityProviderCmd.Flags().String("provider-name", "", "The name of the IdP that you want to describe.")
		cognitoIdp_describeIdentityProviderCmd.Flags().String("user-pool-id", "", "The ID of the user pool that has the IdP that you want to describe..")
		cognitoIdp_describeIdentityProviderCmd.MarkFlagRequired("provider-name")
		cognitoIdp_describeIdentityProviderCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_describeIdentityProviderCmd)
}
