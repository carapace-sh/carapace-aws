package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_updateResourceServerCmd = &cobra.Command{
	Use:   "update-resource-server",
	Short: "Updates the name and scopes of a resource server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_updateResourceServerCmd).Standalone()

	cognitoIdp_updateResourceServerCmd.Flags().String("identifier", "", "A unique resource server identifier for the resource server.")
	cognitoIdp_updateResourceServerCmd.Flags().String("name", "", "The updated name of the resource server.")
	cognitoIdp_updateResourceServerCmd.Flags().String("scopes", "", "An array of updated custom scope names and descriptions that you want to associate with your resource server.")
	cognitoIdp_updateResourceServerCmd.Flags().String("user-pool-id", "", "The ID of the user pool that contains the resource server that you want to update.")
	cognitoIdp_updateResourceServerCmd.MarkFlagRequired("identifier")
	cognitoIdp_updateResourceServerCmd.MarkFlagRequired("name")
	cognitoIdp_updateResourceServerCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_updateResourceServerCmd)
}
