package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_getUserAttributeVerificationCodeCmd = &cobra.Command{
	Use:   "get-user-attribute-verification-code",
	Short: "Given an attribute name, sends a user attribute verification code for the specified attribute name to the currently signed-in user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_getUserAttributeVerificationCodeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_getUserAttributeVerificationCodeCmd).Standalone()

		cognitoIdp_getUserAttributeVerificationCodeCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
		cognitoIdp_getUserAttributeVerificationCodeCmd.Flags().String("attribute-name", "", "The name of the attribute that the user wants to verify, for example `email`.")
		cognitoIdp_getUserAttributeVerificationCodeCmd.Flags().String("client-metadata", "", "A map of custom key-value pairs that you can provide as input for any custom workflows that this action triggers.")
		cognitoIdp_getUserAttributeVerificationCodeCmd.MarkFlagRequired("access-token")
		cognitoIdp_getUserAttributeVerificationCodeCmd.MarkFlagRequired("attribute-name")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_getUserAttributeVerificationCodeCmd)
}
