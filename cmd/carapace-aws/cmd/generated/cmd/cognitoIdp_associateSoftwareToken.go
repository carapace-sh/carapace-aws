package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_associateSoftwareTokenCmd = &cobra.Command{
	Use:   "associate-software-token",
	Short: "Begins setup of time-based one-time password (TOTP) multi-factor authentication (MFA) for a user, with a unique private key that Amazon Cognito generates and returns in the API response.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_associateSoftwareTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_associateSoftwareTokenCmd).Standalone()

		cognitoIdp_associateSoftwareTokenCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
		cognitoIdp_associateSoftwareTokenCmd.Flags().String("session", "", "The session identifier that maintains the state of authentication requests and challenge responses.")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_associateSoftwareTokenCmd)
}
