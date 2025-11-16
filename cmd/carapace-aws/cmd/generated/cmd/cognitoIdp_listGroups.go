package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_listGroupsCmd = &cobra.Command{
	Use:   "list-groups",
	Short: "Given a user pool ID, returns user pool groups and their details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_listGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_listGroupsCmd).Standalone()

		cognitoIdp_listGroupsCmd.Flags().String("limit", "", "The maximum number of groups that you want Amazon Cognito to return in the response.")
		cognitoIdp_listGroupsCmd.Flags().String("next-token", "", "This API operation returns a limited number of results.")
		cognitoIdp_listGroupsCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to list user groups.")
		cognitoIdp_listGroupsCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_listGroupsCmd)
}
