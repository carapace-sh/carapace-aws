package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminListGroupsForUserCmd = &cobra.Command{
	Use:   "admin-list-groups-for-user",
	Short: "Lists the groups that a user belongs to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminListGroupsForUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_adminListGroupsForUserCmd).Standalone()

		cognitoIdp_adminListGroupsForUserCmd.Flags().String("limit", "", "The maximum number of groups that you want Amazon Cognito to return in the response.")
		cognitoIdp_adminListGroupsForUserCmd.Flags().String("next-token", "", "This API operation returns a limited number of results.")
		cognitoIdp_adminListGroupsForUserCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to view a user's groups.")
		cognitoIdp_adminListGroupsForUserCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
		cognitoIdp_adminListGroupsForUserCmd.MarkFlagRequired("user-pool-id")
		cognitoIdp_adminListGroupsForUserCmd.MarkFlagRequired("username")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_adminListGroupsForUserCmd)
}
