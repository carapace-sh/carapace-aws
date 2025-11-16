package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_createAppInstanceBotCmd = &cobra.Command{
	Use:   "create-app-instance-bot",
	Short: "Creates a bot under an Amazon Chime `AppInstance`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_createAppInstanceBotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkIdentity_createAppInstanceBotCmd).Standalone()

		chimeSdkIdentity_createAppInstanceBotCmd.Flags().String("app-instance-arn", "", "The ARN of the `AppInstance` request.")
		chimeSdkIdentity_createAppInstanceBotCmd.Flags().String("client-request-token", "", "The unique ID for the client making the request.")
		chimeSdkIdentity_createAppInstanceBotCmd.Flags().String("configuration", "", "Configuration information about the Amazon Lex V2 V2 bot.")
		chimeSdkIdentity_createAppInstanceBotCmd.Flags().String("metadata", "", "The request metadata.")
		chimeSdkIdentity_createAppInstanceBotCmd.Flags().String("name", "", "The user's name.")
		chimeSdkIdentity_createAppInstanceBotCmd.Flags().String("tags", "", "The tags assigned to the `AppInstanceBot`.")
		chimeSdkIdentity_createAppInstanceBotCmd.MarkFlagRequired("app-instance-arn")
		chimeSdkIdentity_createAppInstanceBotCmd.MarkFlagRequired("client-request-token")
		chimeSdkIdentity_createAppInstanceBotCmd.MarkFlagRequired("configuration")
	})
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_createAppInstanceBotCmd)
}
