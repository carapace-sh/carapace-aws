package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_deleteAppInstanceBotCmd = &cobra.Command{
	Use:   "delete-app-instance-bot",
	Short: "Deletes an `AppInstanceBot`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_deleteAppInstanceBotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkIdentity_deleteAppInstanceBotCmd).Standalone()

		chimeSdkIdentity_deleteAppInstanceBotCmd.Flags().String("app-instance-bot-arn", "", "The ARN of the `AppInstanceBot` being deleted.")
		chimeSdkIdentity_deleteAppInstanceBotCmd.MarkFlagRequired("app-instance-bot-arn")
	})
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_deleteAppInstanceBotCmd)
}
