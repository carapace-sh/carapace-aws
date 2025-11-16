package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_deleteWebAuthnCredentialCmd = &cobra.Command{
	Use:   "delete-web-authn-credential",
	Short: "Deletes a registered passkey, or WebAuthn, authenticator for the currently signed-in user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_deleteWebAuthnCredentialCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_deleteWebAuthnCredentialCmd).Standalone()

		cognitoIdp_deleteWebAuthnCredentialCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
		cognitoIdp_deleteWebAuthnCredentialCmd.Flags().String("credential-id", "", "The unique identifier of the passkey that you want to delete.")
		cognitoIdp_deleteWebAuthnCredentialCmd.MarkFlagRequired("access-token")
		cognitoIdp_deleteWebAuthnCredentialCmd.MarkFlagRequired("credential-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_deleteWebAuthnCredentialCmd)
}
