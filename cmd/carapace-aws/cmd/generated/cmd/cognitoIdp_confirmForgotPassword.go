package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_confirmForgotPasswordCmd = &cobra.Command{
	Use:   "confirm-forgot-password",
	Short: "This public API operation accepts a confirmation code that Amazon Cognito sent to a user and accepts a new password for that user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_confirmForgotPasswordCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_confirmForgotPasswordCmd).Standalone()

		cognitoIdp_confirmForgotPasswordCmd.Flags().String("analytics-metadata", "", "Information that supports analytics outcomes with Amazon Pinpoint, including the user's endpoint ID.")
		cognitoIdp_confirmForgotPasswordCmd.Flags().String("client-id", "", "The ID of the app client where the user wants to reset their password.")
		cognitoIdp_confirmForgotPasswordCmd.Flags().String("client-metadata", "", "A map of custom key-value pairs that you can provide as input for any custom workflows that this action triggers.")
		cognitoIdp_confirmForgotPasswordCmd.Flags().String("confirmation-code", "", "The confirmation code that your user pool delivered when your user requested to reset their password.")
		cognitoIdp_confirmForgotPasswordCmd.Flags().String("password", "", "The new password that your user wants to set.")
		cognitoIdp_confirmForgotPasswordCmd.Flags().String("secret-hash", "", "A keyed-hash message authentication code (HMAC) calculated using the secret key of a user pool client and username plus the client ID in the message.")
		cognitoIdp_confirmForgotPasswordCmd.Flags().String("user-context-data", "", "Contextual data about your user session like the device fingerprint, IP address, or location.")
		cognitoIdp_confirmForgotPasswordCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
		cognitoIdp_confirmForgotPasswordCmd.MarkFlagRequired("client-id")
		cognitoIdp_confirmForgotPasswordCmd.MarkFlagRequired("confirmation-code")
		cognitoIdp_confirmForgotPasswordCmd.MarkFlagRequired("password")
		cognitoIdp_confirmForgotPasswordCmd.MarkFlagRequired("username")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_confirmForgotPasswordCmd)
}
