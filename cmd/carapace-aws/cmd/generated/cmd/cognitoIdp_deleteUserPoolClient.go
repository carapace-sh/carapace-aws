package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_deleteUserPoolClientCmd = &cobra.Command{
	Use:   "delete-user-pool-client",
	Short: "Deletes a user pool app client.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_deleteUserPoolClientCmd).Standalone()

	cognitoIdp_deleteUserPoolClientCmd.Flags().String("client-id", "", "The ID of the user pool app client that you want to delete.")
	cognitoIdp_deleteUserPoolClientCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to delete the client.")
	cognitoIdp_deleteUserPoolClientCmd.MarkFlagRequired("client-id")
	cognitoIdp_deleteUserPoolClientCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_deleteUserPoolClientCmd)
}
