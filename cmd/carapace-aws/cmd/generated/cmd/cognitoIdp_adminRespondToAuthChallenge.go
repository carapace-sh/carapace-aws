package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminRespondToAuthChallengeCmd = &cobra.Command{
	Use:   "admin-respond-to-auth-challenge",
	Short: "Some API operations in a user pool generate a challenge, like a prompt for an MFA code, for device authentication that bypasses MFA, or for a custom authentication challenge.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminRespondToAuthChallengeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_adminRespondToAuthChallengeCmd).Standalone()

		cognitoIdp_adminRespondToAuthChallengeCmd.Flags().String("analytics-metadata", "", "Information that supports analytics outcomes with Amazon Pinpoint, including the user's endpoint ID.")
		cognitoIdp_adminRespondToAuthChallengeCmd.Flags().String("challenge-name", "", "The name of the challenge that you are responding to.")
		cognitoIdp_adminRespondToAuthChallengeCmd.Flags().String("challenge-responses", "", "The responses to the challenge that you received in the previous request.")
		cognitoIdp_adminRespondToAuthChallengeCmd.Flags().String("client-id", "", "The ID of the app client where you initiated sign-in.")
		cognitoIdp_adminRespondToAuthChallengeCmd.Flags().String("client-metadata", "", "A map of custom key-value pairs that you can provide as input for any custom workflows that this action triggers.")
		cognitoIdp_adminRespondToAuthChallengeCmd.Flags().String("context-data", "", "Contextual data about your user session like the device fingerprint, IP address, or location.")
		cognitoIdp_adminRespondToAuthChallengeCmd.Flags().String("session", "", "The session identifier that maintains the state of authentication requests and challenge responses.")
		cognitoIdp_adminRespondToAuthChallengeCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to respond to an authentication challenge.")
		cognitoIdp_adminRespondToAuthChallengeCmd.MarkFlagRequired("challenge-name")
		cognitoIdp_adminRespondToAuthChallengeCmd.MarkFlagRequired("client-id")
		cognitoIdp_adminRespondToAuthChallengeCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_adminRespondToAuthChallengeCmd)
}
