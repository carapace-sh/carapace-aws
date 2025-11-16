package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminRemoveUserFromGroupCmd = &cobra.Command{
	Use:   "admin-remove-user-from-group",
	Short: "Given a username and a group name, removes them from the group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminRemoveUserFromGroupCmd).Standalone()

	cognitoIdp_adminRemoveUserFromGroupCmd.Flags().String("group-name", "", "The name of the group that you want to remove the user from, for example `MyTestGroup`.")
	cognitoIdp_adminRemoveUserFromGroupCmd.Flags().String("user-pool-id", "", "The ID of the user pool that contains the group and the user that you want to remove.")
	cognitoIdp_adminRemoveUserFromGroupCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
	cognitoIdp_adminRemoveUserFromGroupCmd.MarkFlagRequired("group-name")
	cognitoIdp_adminRemoveUserFromGroupCmd.MarkFlagRequired("user-pool-id")
	cognitoIdp_adminRemoveUserFromGroupCmd.MarkFlagRequired("username")
	cognitoIdpCmd.AddCommand(cognitoIdp_adminRemoveUserFromGroupCmd)
}
