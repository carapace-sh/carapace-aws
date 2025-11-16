package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_deleteUserAttributesCmd = &cobra.Command{
	Use:   "delete-user-attributes",
	Short: "Deletes attributes from the currently signed-in user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_deleteUserAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_deleteUserAttributesCmd).Standalone()

		cognitoIdp_deleteUserAttributesCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
		cognitoIdp_deleteUserAttributesCmd.Flags().String("user-attribute-names", "", "An array of strings representing the user attribute names you want to delete.")
		cognitoIdp_deleteUserAttributesCmd.MarkFlagRequired("access-token")
		cognitoIdp_deleteUserAttributesCmd.MarkFlagRequired("user-attribute-names")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_deleteUserAttributesCmd)
}
