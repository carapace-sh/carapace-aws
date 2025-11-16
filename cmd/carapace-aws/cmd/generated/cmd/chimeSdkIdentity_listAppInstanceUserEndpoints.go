package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_listAppInstanceUserEndpointsCmd = &cobra.Command{
	Use:   "list-app-instance-user-endpoints",
	Short: "Lists all the `AppInstanceUserEndpoints` created under a single `AppInstanceUser`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_listAppInstanceUserEndpointsCmd).Standalone()

	chimeSdkIdentity_listAppInstanceUserEndpointsCmd.Flags().String("app-instance-user-arn", "", "The ARN of the `AppInstanceUser`.")
	chimeSdkIdentity_listAppInstanceUserEndpointsCmd.Flags().String("max-results", "", "The maximum number of endpoints that you want to return.")
	chimeSdkIdentity_listAppInstanceUserEndpointsCmd.Flags().String("next-token", "", "The token passed by previous API calls until all requested endpoints are returned.")
	chimeSdkIdentity_listAppInstanceUserEndpointsCmd.MarkFlagRequired("app-instance-user-arn")
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_listAppInstanceUserEndpointsCmd)
}
