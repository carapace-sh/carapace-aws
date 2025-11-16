package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_resendConfirmationCodeCmd = &cobra.Command{
	Use:   "resend-confirmation-code",
	Short: "Resends the code that confirms a new account for a user who has signed up in your user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_resendConfirmationCodeCmd).Standalone()

	cognitoIdp_resendConfirmationCodeCmd.Flags().String("analytics-metadata", "", "Information that supports analytics outcomes with Amazon Pinpoint, including the user's endpoint ID.")
	cognitoIdp_resendConfirmationCodeCmd.Flags().String("client-id", "", "The ID of the user pool app client where the user signed up.")
	cognitoIdp_resendConfirmationCodeCmd.Flags().String("client-metadata", "", "A map of custom key-value pairs that you can provide as input for any custom workflows that this action triggers.")
	cognitoIdp_resendConfirmationCodeCmd.Flags().String("secret-hash", "", "A keyed-hash message authentication code (HMAC) calculated using the secret key of a user pool client and username plus the client ID in the message.")
	cognitoIdp_resendConfirmationCodeCmd.Flags().String("user-context-data", "", "Contextual data about your user session like the device fingerprint, IP address, or location.")
	cognitoIdp_resendConfirmationCodeCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
	cognitoIdp_resendConfirmationCodeCmd.MarkFlagRequired("client-id")
	cognitoIdp_resendConfirmationCodeCmd.MarkFlagRequired("username")
	cognitoIdpCmd.AddCommand(cognitoIdp_resendConfirmationCodeCmd)
}
