package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_updateSlackChannelConfigurationCmd = &cobra.Command{
	Use:   "update-slack-channel-configuration",
	Short: "Updates a Slack channel configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_updateSlackChannelConfigurationCmd).Standalone()

	chatbot_updateSlackChannelConfigurationCmd.Flags().String("chat-configuration-arn", "", "The Amazon Resource Name (ARN) of the SlackChannelConfiguration to update.")
	chatbot_updateSlackChannelConfigurationCmd.Flags().String("guardrail-policy-arns", "", "The list of IAM policy ARNs that are applied as channel guardrails.")
	chatbot_updateSlackChannelConfigurationCmd.Flags().String("iam-role-arn", "", "A user-defined role that AWS Chatbot assumes.")
	chatbot_updateSlackChannelConfigurationCmd.Flags().String("logging-level", "", "Logging levels include `ERROR`, `INFO`, or `NONE`.")
	chatbot_updateSlackChannelConfigurationCmd.Flags().String("slack-channel-id", "", "The ID of the Slack channel.")
	chatbot_updateSlackChannelConfigurationCmd.Flags().String("slack-channel-name", "", "The name of the Slack channel.")
	chatbot_updateSlackChannelConfigurationCmd.Flags().String("sns-topic-arns", "", "The Amazon Resource Names (ARNs) of the SNS topics that deliver notifications to AWS Chatbot.")
	chatbot_updateSlackChannelConfigurationCmd.Flags().String("user-authorization-required", "", "Enables use of a user role requirement in your chat configuration.")
	chatbot_updateSlackChannelConfigurationCmd.MarkFlagRequired("chat-configuration-arn")
	chatbot_updateSlackChannelConfigurationCmd.MarkFlagRequired("slack-channel-id")
	chatbotCmd.AddCommand(chatbot_updateSlackChannelConfigurationCmd)
}
