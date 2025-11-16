package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_updateUserAttributesCmd = &cobra.Command{
	Use:   "update-user-attributes",
	Short: "Updates the currently signed-in user's attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_updateUserAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_updateUserAttributesCmd).Standalone()

		cognitoIdp_updateUserAttributesCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
		cognitoIdp_updateUserAttributesCmd.Flags().String("client-metadata", "", "A map of custom key-value pairs that you can provide as input for any custom workflows that this action initiates.")
		cognitoIdp_updateUserAttributesCmd.Flags().String("user-attributes", "", "An array of name-value pairs representing user attributes.")
		cognitoIdp_updateUserAttributesCmd.MarkFlagRequired("access-token")
		cognitoIdp_updateUserAttributesCmd.MarkFlagRequired("user-attributes")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_updateUserAttributesCmd)
}
