package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_completeWebAuthnRegistrationCmd = &cobra.Command{
	Use:   "complete-web-authn-registration",
	Short: "Completes registration of a passkey authenticator for the currently signed-in user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_completeWebAuthnRegistrationCmd).Standalone()

	cognitoIdp_completeWebAuthnRegistrationCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
	cognitoIdp_completeWebAuthnRegistrationCmd.Flags().String("credential", "", "A [RegistrationResponseJSON](https://www.w3.org/TR/WebAuthn-3/#dictdef-registrationresponsejson) public-key credential response from the user's passkey provider.")
	cognitoIdp_completeWebAuthnRegistrationCmd.MarkFlagRequired("access-token")
	cognitoIdp_completeWebAuthnRegistrationCmd.MarkFlagRequired("credential")
	cognitoIdpCmd.AddCommand(cognitoIdp_completeWebAuthnRegistrationCmd)
}
