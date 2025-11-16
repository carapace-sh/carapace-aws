package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminInitiateAuthCmd = &cobra.Command{
	Use:   "admin-initiate-auth",
	Short: "Starts sign-in for applications with a server-side component, for example a traditional web application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminInitiateAuthCmd).Standalone()

	cognitoIdp_adminInitiateAuthCmd.Flags().String("analytics-metadata", "", "Information that supports analytics outcomes with Amazon Pinpoint, including the user's endpoint ID.")
	cognitoIdp_adminInitiateAuthCmd.Flags().String("auth-flow", "", "The authentication flow that you want to initiate.")
	cognitoIdp_adminInitiateAuthCmd.Flags().String("auth-parameters", "", "The authentication parameters.")
	cognitoIdp_adminInitiateAuthCmd.Flags().String("client-id", "", "The ID of the app client where the user wants to sign in.")
	cognitoIdp_adminInitiateAuthCmd.Flags().String("client-metadata", "", "A map of custom key-value pairs that you can provide as input for certain custom workflows that this action triggers.")
	cognitoIdp_adminInitiateAuthCmd.Flags().String("context-data", "", "Contextual data about your user session like the device fingerprint, IP address, or location.")
	cognitoIdp_adminInitiateAuthCmd.Flags().String("session", "", "The optional session ID from a `ConfirmSignUp` API request.")
	cognitoIdp_adminInitiateAuthCmd.Flags().String("user-pool-id", "", "The ID of the user pool where the user wants to sign in.")
	cognitoIdp_adminInitiateAuthCmd.MarkFlagRequired("auth-flow")
	cognitoIdp_adminInitiateAuthCmd.MarkFlagRequired("client-id")
	cognitoIdp_adminInitiateAuthCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_adminInitiateAuthCmd)
}
