package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_verifyUserAttributeCmd = &cobra.Command{
	Use:   "verify-user-attribute",
	Short: "Submits a verification code for a signed-in user who has added or changed a value of an auto-verified attribute.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_verifyUserAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_verifyUserAttributeCmd).Standalone()

		cognitoIdp_verifyUserAttributeCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
		cognitoIdp_verifyUserAttributeCmd.Flags().String("attribute-name", "", "The name of the attribute that you want to verify.")
		cognitoIdp_verifyUserAttributeCmd.Flags().String("code", "", "The verification code that your user pool sent to the added or changed attribute, for example the user's email address.")
		cognitoIdp_verifyUserAttributeCmd.MarkFlagRequired("access-token")
		cognitoIdp_verifyUserAttributeCmd.MarkFlagRequired("attribute-name")
		cognitoIdp_verifyUserAttributeCmd.MarkFlagRequired("code")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_verifyUserAttributeCmd)
}
