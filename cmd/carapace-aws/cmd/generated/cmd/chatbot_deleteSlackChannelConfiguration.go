package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_deleteSlackChannelConfigurationCmd = &cobra.Command{
	Use:   "delete-slack-channel-configuration",
	Short: "Deletes a Slack channel configuration for AWS Chatbot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_deleteSlackChannelConfigurationCmd).Standalone()

	chatbot_deleteSlackChannelConfigurationCmd.Flags().String("chat-configuration-arn", "", "The Amazon Resource Name (ARN) of the SlackChannelConfiguration to delete.")
	chatbot_deleteSlackChannelConfigurationCmd.MarkFlagRequired("chat-configuration-arn")
	chatbotCmd.AddCommand(chatbot_deleteSlackChannelConfigurationCmd)
}
