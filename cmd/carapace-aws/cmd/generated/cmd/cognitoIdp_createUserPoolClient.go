package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_createUserPoolClientCmd = &cobra.Command{
	Use:   "create-user-pool-client",
	Short: "Creates an app client in a user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_createUserPoolClientCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_createUserPoolClientCmd).Standalone()

		cognitoIdp_createUserPoolClientCmd.Flags().String("access-token-validity", "", "The access token time limit.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("allowed-oauth-flows", "", "The OAuth grant types that you want your app client to generate for clients in managed login authentication.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("allowed-oauth-flows-user-pool-client", "", "Set to `true` to use OAuth 2.0 authorization server features in your app client.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("allowed-oauth-scopes", "", "The OAuth, OpenID Connect (OIDC), and custom scopes that you want to permit your app client to authorize access with.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("analytics-configuration", "", "The user pool analytics configuration for collecting metrics and sending them to your Amazon Pinpoint campaign.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("auth-session-validity", "", "Amazon Cognito creates a session token for each API request in an authentication flow.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("callback-urls", "", "A list of allowed redirect, or callback, URLs for managed login authentication.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("client-name", "", "A friendly name for the app client that you want to create.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("default-redirect-uri", "", "The default redirect URI.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("enable-propagate-additional-user-context-data", "", "When `true`, your application can include additional `UserContextData` in authentication requests.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("enable-token-revocation", "", "Activates or deactivates [token revocation](https://docs.aws.amazon.com/cognito/latest/developerguide/token-revocation.html) in the target app client.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("explicit-auth-flows", "", "The [authentication flows](https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-authentication-flow-methods.html) that you want your user pool client to support.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("generate-secret", "", "When `true`, generates a client secret for the app client.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("id-token-validity", "", "The ID token time limit.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("logout-urls", "", "A list of allowed logout URLs for managed login authentication.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("prevent-user-existence-errors", "", "When `ENABLED`, suppresses messages that might indicate a valid user exists when someone attempts sign-in.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("read-attributes", "", "The list of user attributes that you want your app client to have read access to.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("refresh-token-rotation", "", "The configuration of your app client for refresh token rotation.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("refresh-token-validity", "", "The refresh token time limit.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("supported-identity-providers", "", "A list of provider names for the identity providers (IdPs) that are supported on this client.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("token-validity-units", "", "The units that validity times are represented in.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to create an app client.")
		cognitoIdp_createUserPoolClientCmd.Flags().String("write-attributes", "", "The list of user attributes that you want your app client to have write access to.")
		cognitoIdp_createUserPoolClientCmd.MarkFlagRequired("client-name")
		cognitoIdp_createUserPoolClientCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_createUserPoolClientCmd)
}
