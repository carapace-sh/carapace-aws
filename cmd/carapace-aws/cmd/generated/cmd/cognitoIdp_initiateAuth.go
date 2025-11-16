package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_initiateAuthCmd = &cobra.Command{
	Use:   "initiate-auth",
	Short: "Declares an authentication flow and initiates sign-in for a user in the Amazon Cognito user directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_initiateAuthCmd).Standalone()

	cognitoIdp_initiateAuthCmd.Flags().String("analytics-metadata", "", "Information that supports analytics outcomes with Amazon Pinpoint, including the user's endpoint ID.")
	cognitoIdp_initiateAuthCmd.Flags().String("auth-flow", "", "The authentication flow that you want to initiate.")
	cognitoIdp_initiateAuthCmd.Flags().String("auth-parameters", "", "The authentication parameters.")
	cognitoIdp_initiateAuthCmd.Flags().String("client-id", "", "The ID of the app client that your user wants to sign in to.")
	cognitoIdp_initiateAuthCmd.Flags().String("client-metadata", "", "A map of custom key-value pairs that you can provide as input for certain custom workflows that this action triggers.")
	cognitoIdp_initiateAuthCmd.Flags().String("session", "", "The optional session ID from a `ConfirmSignUp` API request.")
	cognitoIdp_initiateAuthCmd.Flags().String("user-context-data", "", "Contextual data about your user session like the device fingerprint, IP address, or location.")
	cognitoIdp_initiateAuthCmd.MarkFlagRequired("auth-flow")
	cognitoIdp_initiateAuthCmd.MarkFlagRequired("client-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_initiateAuthCmd)
}
