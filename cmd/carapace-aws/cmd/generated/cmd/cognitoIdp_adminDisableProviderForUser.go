package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminDisableProviderForUserCmd = &cobra.Command{
	Use:   "admin-disable-provider-for-user",
	Short: "Prevents the user from signing in with the specified external (SAML or social) identity provider (IdP).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminDisableProviderForUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_adminDisableProviderForUserCmd).Standalone()

		cognitoIdp_adminDisableProviderForUserCmd.Flags().String("user", "", "The user profile that you want to delete a linked identity from.")
		cognitoIdp_adminDisableProviderForUserCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to delete the user's linked identities.")
		cognitoIdp_adminDisableProviderForUserCmd.MarkFlagRequired("user")
		cognitoIdp_adminDisableProviderForUserCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_adminDisableProviderForUserCmd)
}
