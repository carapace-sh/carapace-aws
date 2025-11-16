package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_deleteSlackUserIdentityCmd = &cobra.Command{
	Use:   "delete-slack-user-identity",
	Short: "Deletes a user level permission for a Slack channel configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_deleteSlackUserIdentityCmd).Standalone()

	chatbot_deleteSlackUserIdentityCmd.Flags().String("chat-configuration-arn", "", "The ARN of the SlackChannelConfiguration associated with the user identity to delete.")
	chatbot_deleteSlackUserIdentityCmd.Flags().String("slack-team-id", "", "The ID of the Slack workspace authorized with AWS Chatbot.")
	chatbot_deleteSlackUserIdentityCmd.Flags().String("slack-user-id", "", "The ID of the user in Slack")
	chatbot_deleteSlackUserIdentityCmd.MarkFlagRequired("chat-configuration-arn")
	chatbot_deleteSlackUserIdentityCmd.MarkFlagRequired("slack-team-id")
	chatbot_deleteSlackUserIdentityCmd.MarkFlagRequired("slack-user-id")
	chatbotCmd.AddCommand(chatbot_deleteSlackUserIdentityCmd)
}
