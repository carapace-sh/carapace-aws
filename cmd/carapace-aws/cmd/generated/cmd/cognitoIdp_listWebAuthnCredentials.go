package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_listWebAuthnCredentialsCmd = &cobra.Command{
	Use:   "list-web-authn-credentials",
	Short: "Generates a list of the currently signed-in user's registered passkey, or WebAuthn, credentials.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_listWebAuthnCredentialsCmd).Standalone()

	cognitoIdp_listWebAuthnCredentialsCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
	cognitoIdp_listWebAuthnCredentialsCmd.Flags().String("max-results", "", "The maximum number of the user's passkey credentials that you want to return.")
	cognitoIdp_listWebAuthnCredentialsCmd.Flags().String("next-token", "", "This API operation returns a limited number of results.")
	cognitoIdp_listWebAuthnCredentialsCmd.MarkFlagRequired("access-token")
	cognitoIdpCmd.AddCommand(cognitoIdp_listWebAuthnCredentialsCmd)
}
