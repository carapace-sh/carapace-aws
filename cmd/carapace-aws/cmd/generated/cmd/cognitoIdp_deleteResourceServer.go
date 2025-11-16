package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_deleteResourceServerCmd = &cobra.Command{
	Use:   "delete-resource-server",
	Short: "Deletes a resource server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_deleteResourceServerCmd).Standalone()

	cognitoIdp_deleteResourceServerCmd.Flags().String("identifier", "", "The identifier of the resource server that you want to delete.")
	cognitoIdp_deleteResourceServerCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to delete the resource server.")
	cognitoIdp_deleteResourceServerCmd.MarkFlagRequired("identifier")
	cognitoIdp_deleteResourceServerCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_deleteResourceServerCmd)
}
