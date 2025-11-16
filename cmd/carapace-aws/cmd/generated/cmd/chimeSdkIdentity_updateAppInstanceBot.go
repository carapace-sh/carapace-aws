package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_updateAppInstanceBotCmd = &cobra.Command{
	Use:   "update-app-instance-bot",
	Short: "Updates the name and metadata of an `AppInstanceBot`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_updateAppInstanceBotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkIdentity_updateAppInstanceBotCmd).Standalone()

		chimeSdkIdentity_updateAppInstanceBotCmd.Flags().String("app-instance-bot-arn", "", "The ARN of the `AppInstanceBot`.")
		chimeSdkIdentity_updateAppInstanceBotCmd.Flags().String("configuration", "", "The configuration for the bot update.")
		chimeSdkIdentity_updateAppInstanceBotCmd.Flags().String("metadata", "", "The metadata of the `AppInstanceBot`.")
		chimeSdkIdentity_updateAppInstanceBotCmd.Flags().String("name", "", "The name of the `AppInstanceBot`.")
		chimeSdkIdentity_updateAppInstanceBotCmd.MarkFlagRequired("app-instance-bot-arn")
		chimeSdkIdentity_updateAppInstanceBotCmd.MarkFlagRequired("metadata")
		chimeSdkIdentity_updateAppInstanceBotCmd.MarkFlagRequired("name")
	})
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_updateAppInstanceBotCmd)
}
