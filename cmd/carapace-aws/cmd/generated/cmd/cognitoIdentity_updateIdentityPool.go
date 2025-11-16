package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_updateIdentityPoolCmd = &cobra.Command{
	Use:   "update-identity-pool",
	Short: "Updates the configuration of an identity pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_updateIdentityPoolCmd).Standalone()

	cognitoIdentity_updateIdentityPoolCmd.Flags().String("allow-classic-flow", "", "Enables or disables the Basic (Classic) authentication flow.")
	cognitoIdentity_updateIdentityPoolCmd.Flags().String("allow-unauthenticated-identities", "", "TRUE if the identity pool supports unauthenticated logins.")
	cognitoIdentity_updateIdentityPoolCmd.Flags().String("cognito-identity-providers", "", "A list representing an Amazon Cognito user pool and its client ID.")
	cognitoIdentity_updateIdentityPoolCmd.Flags().String("developer-provider-name", "", "The \"domain\" by which Cognito will refer to your users.")
	cognitoIdentity_updateIdentityPoolCmd.Flags().String("identity-pool-id", "", "An identity pool ID in the format REGION:GUID.")
	cognitoIdentity_updateIdentityPoolCmd.Flags().String("identity-pool-name", "", "A string that you provide.")
	cognitoIdentity_updateIdentityPoolCmd.Flags().String("identity-pool-tags", "", "The tags that are assigned to the identity pool.")
	cognitoIdentity_updateIdentityPoolCmd.Flags().String("open-id-connect-provider-arns", "", "The ARNs of the OpenID Connect providers.")
	cognitoIdentity_updateIdentityPoolCmd.Flags().String("saml-provider-arns", "", "An array of Amazon Resource Names (ARNs) of the SAML provider for your identity pool.")
	cognitoIdentity_updateIdentityPoolCmd.Flags().String("supported-login-providers", "", "Optional key:value pairs mapping provider names to provider app IDs.")
	cognitoIdentity_updateIdentityPoolCmd.MarkFlagRequired("allow-unauthenticated-identities")
	cognitoIdentity_updateIdentityPoolCmd.MarkFlagRequired("identity-pool-id")
	cognitoIdentity_updateIdentityPoolCmd.MarkFlagRequired("identity-pool-name")
	cognitoIdentityCmd.AddCommand(cognitoIdentity_updateIdentityPoolCmd)
}
