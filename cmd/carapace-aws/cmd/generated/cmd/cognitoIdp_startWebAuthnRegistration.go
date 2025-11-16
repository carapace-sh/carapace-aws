package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_startWebAuthnRegistrationCmd = &cobra.Command{
	Use:   "start-web-authn-registration",
	Short: "Requests credential creation options from your user pool for the currently signed-in user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_startWebAuthnRegistrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_startWebAuthnRegistrationCmd).Standalone()

		cognitoIdp_startWebAuthnRegistrationCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
		cognitoIdp_startWebAuthnRegistrationCmd.MarkFlagRequired("access-token")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_startWebAuthnRegistrationCmd)
}
