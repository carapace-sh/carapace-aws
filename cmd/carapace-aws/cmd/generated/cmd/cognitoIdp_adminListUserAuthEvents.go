package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminListUserAuthEventsCmd = &cobra.Command{
	Use:   "admin-list-user-auth-events",
	Short: "Requests a history of user activity and any risks detected as part of Amazon Cognito threat protection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminListUserAuthEventsCmd).Standalone()

	cognitoIdp_adminListUserAuthEventsCmd.Flags().String("max-results", "", "The maximum number of authentication events to return.")
	cognitoIdp_adminListUserAuthEventsCmd.Flags().String("next-token", "", "This API operation returns a limited number of results.")
	cognitoIdp_adminListUserAuthEventsCmd.Flags().String("user-pool-id", "", "The Id of the user pool that contains the user profile with the logged events.")
	cognitoIdp_adminListUserAuthEventsCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
	cognitoIdp_adminListUserAuthEventsCmd.MarkFlagRequired("user-pool-id")
	cognitoIdp_adminListUserAuthEventsCmd.MarkFlagRequired("username")
	cognitoIdpCmd.AddCommand(cognitoIdp_adminListUserAuthEventsCmd)
}
