package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_listUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "Given a user pool ID, returns a list of users and their basic details in a user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_listUsersCmd).Standalone()

	cognitoIdp_listUsersCmd.Flags().String("attributes-to-get", "", "A JSON array of user attribute names, for example `given_name`, that you want Amazon Cognito to include in the response for each user.")
	cognitoIdp_listUsersCmd.Flags().String("filter", "", "A filter string of the form `\"AttributeName Filter-Type \"AttributeValue\"`.")
	cognitoIdp_listUsersCmd.Flags().String("limit", "", "The maximum number of users that you want Amazon Cognito to return in the response.")
	cognitoIdp_listUsersCmd.Flags().String("pagination-token", "", "This API operation returns a limited number of results.")
	cognitoIdp_listUsersCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to display or search for users.")
	cognitoIdp_listUsersCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_listUsersCmd)
}
