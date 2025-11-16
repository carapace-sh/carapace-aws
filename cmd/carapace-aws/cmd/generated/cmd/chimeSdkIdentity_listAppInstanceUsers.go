package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_listAppInstanceUsersCmd = &cobra.Command{
	Use:   "list-app-instance-users",
	Short: "List all `AppInstanceUsers` created under a single `AppInstance`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_listAppInstanceUsersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkIdentity_listAppInstanceUsersCmd).Standalone()

		chimeSdkIdentity_listAppInstanceUsersCmd.Flags().String("app-instance-arn", "", "The ARN of the `AppInstance`.")
		chimeSdkIdentity_listAppInstanceUsersCmd.Flags().String("max-results", "", "The maximum number of requests that you want returned.")
		chimeSdkIdentity_listAppInstanceUsersCmd.Flags().String("next-token", "", "The token passed by previous API calls until all requested users are returned.")
		chimeSdkIdentity_listAppInstanceUsersCmd.MarkFlagRequired("app-instance-arn")
	})
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_listAppInstanceUsersCmd)
}
