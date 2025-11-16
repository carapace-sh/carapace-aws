package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_signUpCmd = &cobra.Command{
	Use:   "sign-up",
	Short: "Registers a user with an app client and requests a user name, password, and user attributes in the user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_signUpCmd).Standalone()

	cognitoIdp_signUpCmd.Flags().String("analytics-metadata", "", "Information that supports analytics outcomes with Amazon Pinpoint, including the user's endpoint ID.")
	cognitoIdp_signUpCmd.Flags().String("client-id", "", "The ID of the app client where the user wants to sign up.")
	cognitoIdp_signUpCmd.Flags().String("client-metadata", "", "A map of custom key-value pairs that you can provide as input for any custom workflows that this action triggers.")
	cognitoIdp_signUpCmd.Flags().String("password", "", "The user's proposed password.")
	cognitoIdp_signUpCmd.Flags().String("secret-hash", "", "A keyed-hash message authentication code (HMAC) calculated using the secret key of a user pool client and username plus the client ID in the message.")
	cognitoIdp_signUpCmd.Flags().String("user-attributes", "", "An array of name-value pairs representing user attributes.")
	cognitoIdp_signUpCmd.Flags().String("user-context-data", "", "Contextual data about your user session like the device fingerprint, IP address, or location.")
	cognitoIdp_signUpCmd.Flags().String("username", "", "The username of the user that you want to sign up.")
	cognitoIdp_signUpCmd.Flags().String("validation-data", "", "Temporary user attributes that contribute to the outcomes of your pre sign-up Lambda trigger.")
	cognitoIdp_signUpCmd.MarkFlagRequired("client-id")
	cognitoIdp_signUpCmd.MarkFlagRequired("username")
	cognitoIdpCmd.AddCommand(cognitoIdp_signUpCmd)
}
