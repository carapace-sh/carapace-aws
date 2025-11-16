package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_describeSlackUserIdentitiesCmd = &cobra.Command{
	Use:   "describe-slack-user-identities",
	Short: "Lists all Slack user identities with a mapped role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_describeSlackUserIdentitiesCmd).Standalone()

	chatbot_describeSlackUserIdentitiesCmd.Flags().String("chat-configuration-arn", "", "The Amazon Resource Name (ARN) of the SlackChannelConfiguration associated with the user identities to describe.")
	chatbot_describeSlackUserIdentitiesCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
	chatbot_describeSlackUserIdentitiesCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
	chatbotCmd.AddCommand(chatbot_describeSlackUserIdentitiesCmd)
}
