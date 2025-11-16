package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminUpdateUserAttributesCmd = &cobra.Command{
	Use:   "admin-update-user-attributes",
	Short: "Updates the specified user's attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminUpdateUserAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_adminUpdateUserAttributesCmd).Standalone()

		cognitoIdp_adminUpdateUserAttributesCmd.Flags().String("client-metadata", "", "A map of custom key-value pairs that you can provide as input for any custom workflows that this action triggers.")
		cognitoIdp_adminUpdateUserAttributesCmd.Flags().String("user-attributes", "", "An array of name-value pairs representing user attributes.")
		cognitoIdp_adminUpdateUserAttributesCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to update user attributes.")
		cognitoIdp_adminUpdateUserAttributesCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
		cognitoIdp_adminUpdateUserAttributesCmd.MarkFlagRequired("user-attributes")
		cognitoIdp_adminUpdateUserAttributesCmd.MarkFlagRequired("user-pool-id")
		cognitoIdp_adminUpdateUserAttributesCmd.MarkFlagRequired("username")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_adminUpdateUserAttributesCmd)
}
