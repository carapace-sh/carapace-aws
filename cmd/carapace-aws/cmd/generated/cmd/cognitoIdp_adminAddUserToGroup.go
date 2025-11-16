package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminAddUserToGroupCmd = &cobra.Command{
	Use:   "admin-add-user-to-group",
	Short: "Adds a user to a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminAddUserToGroupCmd).Standalone()

	cognitoIdp_adminAddUserToGroupCmd.Flags().String("group-name", "", "The name of the group that you want to add your user to.")
	cognitoIdp_adminAddUserToGroupCmd.Flags().String("user-pool-id", "", "The ID of the user pool that contains the group that you want to add the user to.")
	cognitoIdp_adminAddUserToGroupCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
	cognitoIdp_adminAddUserToGroupCmd.MarkFlagRequired("group-name")
	cognitoIdp_adminAddUserToGroupCmd.MarkFlagRequired("user-pool-id")
	cognitoIdp_adminAddUserToGroupCmd.MarkFlagRequired("username")
	cognitoIdpCmd.AddCommand(cognitoIdp_adminAddUserToGroupCmd)
}
