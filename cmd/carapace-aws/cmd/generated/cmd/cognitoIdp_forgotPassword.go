package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_forgotPasswordCmd = &cobra.Command{
	Use:   "forgot-password",
	Short: "Sends a password-reset confirmation code to the email address or phone number of the requested username.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_forgotPasswordCmd).Standalone()

	cognitoIdp_forgotPasswordCmd.Flags().String("analytics-metadata", "", "Information that supports analytics outcomes with Amazon Pinpoint, including the user's endpoint ID.")
	cognitoIdp_forgotPasswordCmd.Flags().String("client-id", "", "The ID of the user pool app client associated with the current signed-in user.")
	cognitoIdp_forgotPasswordCmd.Flags().String("client-metadata", "", "A map of custom key-value pairs that you can provide as input for any custom workflows that this action triggers.")
	cognitoIdp_forgotPasswordCmd.Flags().String("secret-hash", "", "A keyed-hash message authentication code (HMAC) calculated using the secret key of a user pool client and username plus the client ID in the message.")
	cognitoIdp_forgotPasswordCmd.Flags().String("user-context-data", "", "Contextual data about your user session like the device fingerprint, IP address, or location.")
	cognitoIdp_forgotPasswordCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
	cognitoIdp_forgotPasswordCmd.MarkFlagRequired("client-id")
	cognitoIdp_forgotPasswordCmd.MarkFlagRequired("username")
	cognitoIdpCmd.AddCommand(cognitoIdp_forgotPasswordCmd)
}
