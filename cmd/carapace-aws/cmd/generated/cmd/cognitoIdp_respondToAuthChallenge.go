package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_respondToAuthChallengeCmd = &cobra.Command{
	Use:   "respond-to-auth-challenge",
	Short: "Some API operations in a user pool generate a challenge, like a prompt for an MFA code, for device authentication that bypasses MFA, or for a custom authentication challenge.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_respondToAuthChallengeCmd).Standalone()

	cognitoIdp_respondToAuthChallengeCmd.Flags().String("analytics-metadata", "", "Information that supports analytics outcomes with Amazon Pinpoint, including the user's endpoint ID.")
	cognitoIdp_respondToAuthChallengeCmd.Flags().String("challenge-name", "", "The name of the challenge that you are responding to.")
	cognitoIdp_respondToAuthChallengeCmd.Flags().String("challenge-responses", "", "The responses to the challenge that you received in the previous request.")
	cognitoIdp_respondToAuthChallengeCmd.Flags().String("client-id", "", "The ID of the app client where the user is signing in.")
	cognitoIdp_respondToAuthChallengeCmd.Flags().String("client-metadata", "", "A map of custom key-value pairs that you can provide as input for any custom workflows that this action triggers.")
	cognitoIdp_respondToAuthChallengeCmd.Flags().String("session", "", "The session identifier that maintains the state of authentication requests and challenge responses.")
	cognitoIdp_respondToAuthChallengeCmd.Flags().String("user-context-data", "", "Contextual data about your user session like the device fingerprint, IP address, or location.")
	cognitoIdp_respondToAuthChallengeCmd.MarkFlagRequired("challenge-name")
	cognitoIdp_respondToAuthChallengeCmd.MarkFlagRequired("client-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_respondToAuthChallengeCmd)
}
