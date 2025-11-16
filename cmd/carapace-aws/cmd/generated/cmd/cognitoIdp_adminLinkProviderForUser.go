package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminLinkProviderForUserCmd = &cobra.Command{
	Use:   "admin-link-provider-for-user",
	Short: "Links an existing user account in a user pool, or `DestinationUser`, to an identity from an external IdP, or `SourceUser`, based on a specified attribute name and value from the external IdP.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminLinkProviderForUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_adminLinkProviderForUserCmd).Standalone()

		cognitoIdp_adminLinkProviderForUserCmd.Flags().String("destination-user", "", "The existing user in the user pool that you want to assign to the external IdP user account.")
		cognitoIdp_adminLinkProviderForUserCmd.Flags().String("source-user", "", "An external IdP account for a user who doesn't exist yet in the user pool.")
		cognitoIdp_adminLinkProviderForUserCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to link a federated identity.")
		cognitoIdp_adminLinkProviderForUserCmd.MarkFlagRequired("destination-user")
		cognitoIdp_adminLinkProviderForUserCmd.MarkFlagRequired("source-user")
		cognitoIdp_adminLinkProviderForUserCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_adminLinkProviderForUserCmd)
}
