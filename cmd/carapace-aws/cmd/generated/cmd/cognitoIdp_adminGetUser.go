package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminGetUserCmd = &cobra.Command{
	Use:   "admin-get-user",
	Short: "Given a username, returns details about a user profile in a user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminGetUserCmd).Standalone()

	cognitoIdp_adminGetUserCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to get information about the user.")
	cognitoIdp_adminGetUserCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
	cognitoIdp_adminGetUserCmd.MarkFlagRequired("user-pool-id")
	cognitoIdp_adminGetUserCmd.MarkFlagRequired("username")
	cognitoIdpCmd.AddCommand(cognitoIdp_adminGetUserCmd)
}
