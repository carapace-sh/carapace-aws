package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminDeleteUserAttributesCmd = &cobra.Command{
	Use:   "admin-delete-user-attributes",
	Short: "Deletes attribute values from a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminDeleteUserAttributesCmd).Standalone()

	cognitoIdp_adminDeleteUserAttributesCmd.Flags().String("user-attribute-names", "", "An array of strings representing the user attribute names you want to delete.")
	cognitoIdp_adminDeleteUserAttributesCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to delete user attributes.")
	cognitoIdp_adminDeleteUserAttributesCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
	cognitoIdp_adminDeleteUserAttributesCmd.MarkFlagRequired("user-attribute-names")
	cognitoIdp_adminDeleteUserAttributesCmd.MarkFlagRequired("user-pool-id")
	cognitoIdp_adminDeleteUserAttributesCmd.MarkFlagRequired("username")
	cognitoIdpCmd.AddCommand(cognitoIdp_adminDeleteUserAttributesCmd)
}
