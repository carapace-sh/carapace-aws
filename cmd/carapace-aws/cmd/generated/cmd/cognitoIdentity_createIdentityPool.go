package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_createIdentityPoolCmd = &cobra.Command{
	Use:   "create-identity-pool",
	Short: "Creates a new identity pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_createIdentityPoolCmd).Standalone()

	cognitoIdentity_createIdentityPoolCmd.Flags().String("allow-classic-flow", "", "Enables or disables the Basic (Classic) authentication flow.")
	cognitoIdentity_createIdentityPoolCmd.Flags().String("allow-unauthenticated-identities", "", "TRUE if the identity pool supports unauthenticated logins.")
	cognitoIdentity_createIdentityPoolCmd.Flags().String("cognito-identity-providers", "", "An array of Amazon Cognito user pools and their client IDs.")
	cognitoIdentity_createIdentityPoolCmd.Flags().String("developer-provider-name", "", "The \"domain\" by which Cognito will refer to your users.")
	cognitoIdentity_createIdentityPoolCmd.Flags().String("identity-pool-name", "", "A string that you provide.")
	cognitoIdentity_createIdentityPoolCmd.Flags().String("identity-pool-tags", "", "Tags to assign to the identity pool.")
	cognitoIdentity_createIdentityPoolCmd.Flags().String("open-id-connect-provider-arns", "", "The Amazon Resource Names (ARN) of the OpenID Connect providers.")
	cognitoIdentity_createIdentityPoolCmd.Flags().String("saml-provider-arns", "", "An array of Amazon Resource Names (ARNs) of the SAML provider for your identity pool.")
	cognitoIdentity_createIdentityPoolCmd.Flags().String("supported-login-providers", "", "Optional key:value pairs mapping provider names to provider app IDs.")
	cognitoIdentity_createIdentityPoolCmd.MarkFlagRequired("allow-unauthenticated-identities")
	cognitoIdentity_createIdentityPoolCmd.MarkFlagRequired("identity-pool-name")
	cognitoIdentityCmd.AddCommand(cognitoIdentity_createIdentityPoolCmd)
}
