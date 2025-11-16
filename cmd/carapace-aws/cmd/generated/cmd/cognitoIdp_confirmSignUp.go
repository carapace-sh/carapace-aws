package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_confirmSignUpCmd = &cobra.Command{
	Use:   "confirm-sign-up",
	Short: "Confirms the account of a new user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_confirmSignUpCmd).Standalone()

	cognitoIdp_confirmSignUpCmd.Flags().String("analytics-metadata", "", "Information that supports analytics outcomes with Amazon Pinpoint, including the user's endpoint ID.")
	cognitoIdp_confirmSignUpCmd.Flags().String("client-id", "", "The ID of the app client associated with the user pool.")
	cognitoIdp_confirmSignUpCmd.Flags().String("client-metadata", "", "A map of custom key-value pairs that you can provide as input for any custom workflows that this action triggers.")
	cognitoIdp_confirmSignUpCmd.Flags().String("confirmation-code", "", "The confirmation code that your user pool sent in response to the `SignUp` request.")
	cognitoIdp_confirmSignUpCmd.Flags().String("force-alias-creation", "", "When `true`, forces user confirmation despite any existing aliases.")
	cognitoIdp_confirmSignUpCmd.Flags().String("secret-hash", "", "A keyed-hash message authentication code (HMAC) calculated using the secret key of a user pool client and username plus the client ID in the message.")
	cognitoIdp_confirmSignUpCmd.Flags().String("session", "", "The optional session ID from a `SignUp` API request.")
	cognitoIdp_confirmSignUpCmd.Flags().String("user-context-data", "", "Contextual data about your user session like the device fingerprint, IP address, or location.")
	cognitoIdp_confirmSignUpCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
	cognitoIdp_confirmSignUpCmd.MarkFlagRequired("client-id")
	cognitoIdp_confirmSignUpCmd.MarkFlagRequired("confirmation-code")
	cognitoIdp_confirmSignUpCmd.MarkFlagRequired("username")
	cognitoIdpCmd.AddCommand(cognitoIdp_confirmSignUpCmd)
}
