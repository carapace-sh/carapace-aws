package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_verifySoftwareTokenCmd = &cobra.Command{
	Use:   "verify-software-token",
	Short: "Registers the current user's time-based one-time password (TOTP) authenticator with a code generated in their authenticator app from a private key that's supplied by your user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_verifySoftwareTokenCmd).Standalone()

	cognitoIdp_verifySoftwareTokenCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
	cognitoIdp_verifySoftwareTokenCmd.Flags().String("friendly-device-name", "", "A friendly name for the device that's running the TOTP authenticator.")
	cognitoIdp_verifySoftwareTokenCmd.Flags().String("session", "", "The session ID from an `AssociateSoftwareToken` request.")
	cognitoIdp_verifySoftwareTokenCmd.Flags().String("user-code", "", "A TOTP that the user generated in their configured authenticator app.")
	cognitoIdp_verifySoftwareTokenCmd.MarkFlagRequired("user-code")
	cognitoIdpCmd.AddCommand(cognitoIdp_verifySoftwareTokenCmd)
}
