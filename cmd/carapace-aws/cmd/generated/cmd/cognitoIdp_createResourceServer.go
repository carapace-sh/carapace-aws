package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_createResourceServerCmd = &cobra.Command{
	Use:   "create-resource-server",
	Short: "Creates a new OAuth2.0 resource server and defines custom scopes within it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_createResourceServerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_createResourceServerCmd).Standalone()

		cognitoIdp_createResourceServerCmd.Flags().String("identifier", "", "A unique resource server identifier for the resource server.")
		cognitoIdp_createResourceServerCmd.Flags().String("name", "", "A friendly name for the resource server.")
		cognitoIdp_createResourceServerCmd.Flags().String("scopes", "", "A list of custom scopes.")
		cognitoIdp_createResourceServerCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to create a resource server.")
		cognitoIdp_createResourceServerCmd.MarkFlagRequired("identifier")
		cognitoIdp_createResourceServerCmd.MarkFlagRequired("name")
		cognitoIdp_createResourceServerCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_createResourceServerCmd)
}
