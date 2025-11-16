package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_updateUserPoolClientCmd = &cobra.Command{
	Use:   "update-user-pool-client",
	Short: "Given a user pool app client ID, updates the configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_updateUserPoolClientCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_updateUserPoolClientCmd).Standalone()

		cognitoIdp_updateUserPoolClientCmd.Flags().String("access-token-validity", "", "The access token time limit.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("allowed-oauth-flows", "", "The OAuth grant types that you want your app client to generate.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("allowed-oauth-flows-user-pool-client", "", "Set to `true` to use OAuth 2.0 authorization server features in your app client.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("allowed-oauth-scopes", "", "The OAuth, OpenID Connect (OIDC), and custom scopes that you want to permit your app client to authorize access with.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("analytics-configuration", "", "The user pool analytics configuration for collecting metrics and sending them to your Amazon Pinpoint campaign.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("auth-session-validity", "", "Amazon Cognito creates a session token for each API request in an authentication flow.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("callback-urls", "", "A list of allowed redirect, or callback, URLs for managed login authentication.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("client-id", "", "The ID of the app client that you want to update.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("client-name", "", "A friendly name for the app client.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("default-redirect-uri", "", "The default redirect URI.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("enable-propagate-additional-user-context-data", "", "When `true`, your application can include additional `UserContextData` in authentication requests.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("enable-token-revocation", "", "Activates or deactivates [token revocation](https://docs.aws.amazon.com/cognito/latest/developerguide/token-revocation.html) in the target app client.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("explicit-auth-flows", "", "The [authentication flows](https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-authentication-flow-methods.html) that you want your user pool client to support.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("id-token-validity", "", "The ID token time limit.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("logout-urls", "", "A list of allowed logout URLs for managed login authentication.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("prevent-user-existence-errors", "", "When `ENABLED`, suppresses messages that might indicate a valid user exists when someone attempts sign-in.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("read-attributes", "", "The list of user attributes that you want your app client to have read access to.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("refresh-token-rotation", "", "The configuration of your app client for refresh token rotation.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("refresh-token-validity", "", "The refresh token time limit.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("supported-identity-providers", "", "A list of provider names for the identity providers (IdPs) that are supported on this client.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("token-validity-units", "", "The units that validity times are represented in.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to update the app client.")
		cognitoIdp_updateUserPoolClientCmd.Flags().String("write-attributes", "", "The list of user attributes that you want your app client to have write access to.")
		cognitoIdp_updateUserPoolClientCmd.MarkFlagRequired("client-id")
		cognitoIdp_updateUserPoolClientCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_updateUserPoolClientCmd)
}
