package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_getTokensFromRefreshTokenCmd = &cobra.Command{
	Use:   "get-tokens-from-refresh-token",
	Short: "Given a refresh token, issues new ID, access, and optionally refresh tokens for the user who owns the submitted token.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_getTokensFromRefreshTokenCmd).Standalone()

	cognitoIdp_getTokensFromRefreshTokenCmd.Flags().String("client-id", "", "The app client that issued the refresh token to the user who wants to request new tokens.")
	cognitoIdp_getTokensFromRefreshTokenCmd.Flags().String("client-metadata", "", "A map of custom key-value pairs that you can provide as input for certain custom workflows that this action triggers.")
	cognitoIdp_getTokensFromRefreshTokenCmd.Flags().String("client-secret", "", "The client secret of the requested app client, if the client has a secret.")
	cognitoIdp_getTokensFromRefreshTokenCmd.Flags().String("device-key", "", "When you enable device remembering, Amazon Cognito issues a device key that you can use for device authentication that bypasses multi-factor authentication (MFA).")
	cognitoIdp_getTokensFromRefreshTokenCmd.Flags().String("refresh-token", "", "A valid refresh token that can authorize the request for new tokens.")
	cognitoIdp_getTokensFromRefreshTokenCmd.MarkFlagRequired("client-id")
	cognitoIdp_getTokensFromRefreshTokenCmd.MarkFlagRequired("refresh-token")
	cognitoIdpCmd.AddCommand(cognitoIdp_getTokensFromRefreshTokenCmd)
}
