package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_listAppInstanceBotsCmd = &cobra.Command{
	Use:   "list-app-instance-bots",
	Short: "Lists all `AppInstanceBots` created under a single `AppInstance`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_listAppInstanceBotsCmd).Standalone()

	chimeSdkIdentity_listAppInstanceBotsCmd.Flags().String("app-instance-arn", "", "The ARN of the `AppInstance`.")
	chimeSdkIdentity_listAppInstanceBotsCmd.Flags().String("max-results", "", "The maximum number of requests to return.")
	chimeSdkIdentity_listAppInstanceBotsCmd.Flags().String("next-token", "", "The token passed by previous API calls until all requested bots are returned.")
	chimeSdkIdentity_listAppInstanceBotsCmd.MarkFlagRequired("app-instance-arn")
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_listAppInstanceBotsCmd)
}
