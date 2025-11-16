package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_getOpenIdTokenForDeveloperIdentityCmd = &cobra.Command{
	Use:   "get-open-id-token-for-developer-identity",
	Short: "Registers (or retrieves) a Cognito `IdentityId` and an OpenID Connect token for a user authenticated by your backend authentication process.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_getOpenIdTokenForDeveloperIdentityCmd).Standalone()

	cognitoIdentity_getOpenIdTokenForDeveloperIdentityCmd.Flags().String("identity-id", "", "A unique identifier in the format REGION:GUID.")
	cognitoIdentity_getOpenIdTokenForDeveloperIdentityCmd.Flags().String("identity-pool-id", "", "An identity pool ID in the format REGION:GUID.")
	cognitoIdentity_getOpenIdTokenForDeveloperIdentityCmd.Flags().String("logins", "", "A set of optional name-value pairs that map provider names to provider tokens.")
	cognitoIdentity_getOpenIdTokenForDeveloperIdentityCmd.Flags().String("principal-tags", "", "Use this operation to configure attribute mappings for custom providers.")
	cognitoIdentity_getOpenIdTokenForDeveloperIdentityCmd.Flags().String("token-duration", "", "The expiration time of the token, in seconds.")
	cognitoIdentity_getOpenIdTokenForDeveloperIdentityCmd.MarkFlagRequired("identity-pool-id")
	cognitoIdentity_getOpenIdTokenForDeveloperIdentityCmd.MarkFlagRequired("logins")
	cognitoIdentityCmd.AddCommand(cognitoIdentity_getOpenIdTokenForDeveloperIdentityCmd)
}
