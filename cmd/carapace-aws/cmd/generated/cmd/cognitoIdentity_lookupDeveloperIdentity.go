package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_lookupDeveloperIdentityCmd = &cobra.Command{
	Use:   "lookup-developer-identity",
	Short: "Retrieves the `IdentityID` associated with a `DeveloperUserIdentifier` or the list of `DeveloperUserIdentifier` values associated with an `IdentityId` for an existing identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_lookupDeveloperIdentityCmd).Standalone()

	cognitoIdentity_lookupDeveloperIdentityCmd.Flags().String("developer-user-identifier", "", "A unique ID used by your backend authentication process to identify a user.")
	cognitoIdentity_lookupDeveloperIdentityCmd.Flags().String("identity-id", "", "A unique identifier in the format REGION:GUID.")
	cognitoIdentity_lookupDeveloperIdentityCmd.Flags().String("identity-pool-id", "", "An identity pool ID in the format REGION:GUID.")
	cognitoIdentity_lookupDeveloperIdentityCmd.Flags().String("max-results", "", "The maximum number of identities to return.")
	cognitoIdentity_lookupDeveloperIdentityCmd.Flags().String("next-token", "", "A pagination token.")
	cognitoIdentity_lookupDeveloperIdentityCmd.MarkFlagRequired("identity-pool-id")
	cognitoIdentityCmd.AddCommand(cognitoIdentity_lookupDeveloperIdentityCmd)
}
