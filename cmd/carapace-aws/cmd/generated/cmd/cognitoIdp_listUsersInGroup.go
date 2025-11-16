package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_listUsersInGroupCmd = &cobra.Command{
	Use:   "list-users-in-group",
	Short: "Given a user pool ID and a group name, returns a list of users in the group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_listUsersInGroupCmd).Standalone()

	cognitoIdp_listUsersInGroupCmd.Flags().String("group-name", "", "The name of the group that you want to query for user membership.")
	cognitoIdp_listUsersInGroupCmd.Flags().String("limit", "", "The maximum number of groups that you want Amazon Cognito to return in the response.")
	cognitoIdp_listUsersInGroupCmd.Flags().String("next-token", "", "This API operation returns a limited number of results.")
	cognitoIdp_listUsersInGroupCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to view the membership of the requested group.")
	cognitoIdp_listUsersInGroupCmd.MarkFlagRequired("group-name")
	cognitoIdp_listUsersInGroupCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_listUsersInGroupCmd)
}
