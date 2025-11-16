package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminDeleteUserCmd = &cobra.Command{
	Use:   "admin-delete-user",
	Short: "Deletes a user profile in your user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminDeleteUserCmd).Standalone()

	cognitoIdp_adminDeleteUserCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to delete the user.")
	cognitoIdp_adminDeleteUserCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
	cognitoIdp_adminDeleteUserCmd.MarkFlagRequired("user-pool-id")
	cognitoIdp_adminDeleteUserCmd.MarkFlagRequired("username")
	cognitoIdpCmd.AddCommand(cognitoIdp_adminDeleteUserCmd)
}
