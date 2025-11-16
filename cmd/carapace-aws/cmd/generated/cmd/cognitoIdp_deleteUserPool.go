package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_deleteUserPoolCmd = &cobra.Command{
	Use:   "delete-user-pool",
	Short: "Deletes a user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_deleteUserPoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_deleteUserPoolCmd).Standalone()

		cognitoIdp_deleteUserPoolCmd.Flags().String("user-pool-id", "", "The ID of the user pool that you want to delete.")
		cognitoIdp_deleteUserPoolCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_deleteUserPoolCmd)
}
