package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_updateMicrosoftTeamsChannelConfigurationCmd = &cobra.Command{
	Use:   "update-microsoft-teams-channel-configuration",
	Short: "Updates an Microsoft Teams channel configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_updateMicrosoftTeamsChannelConfigurationCmd).Standalone()

	chatbot_updateMicrosoftTeamsChannelConfigurationCmd.Flags().String("channel-id", "", "The ID of the Microsoft Teams channel.")
	chatbot_updateMicrosoftTeamsChannelConfigurationCmd.Flags().String("channel-name", "", "The name of the Microsoft Teams channel.")
	chatbot_updateMicrosoftTeamsChannelConfigurationCmd.Flags().String("chat-configuration-arn", "", "The Amazon Resource Name (ARN) of the TeamsChannelConfiguration to update.")
	chatbot_updateMicrosoftTeamsChannelConfigurationCmd.Flags().String("guardrail-policy-arns", "", "The list of IAM policy ARNs that are applied as channel guardrails.")
	chatbot_updateMicrosoftTeamsChannelConfigurationCmd.Flags().String("iam-role-arn", "", "A user-defined role that AWS Chatbot assumes.")
	chatbot_updateMicrosoftTeamsChannelConfigurationCmd.Flags().String("logging-level", "", "Logging levels include `ERROR`, `INFO`, or `NONE`.")
	chatbot_updateMicrosoftTeamsChannelConfigurationCmd.Flags().String("sns-topic-arns", "", "The Amazon Resource Names (ARNs) of the SNS topics that deliver notifications to AWS Chatbot.")
	chatbot_updateMicrosoftTeamsChannelConfigurationCmd.Flags().String("user-authorization-required", "", "Enables use of a user role requirement in your chat configuration.")
	chatbot_updateMicrosoftTeamsChannelConfigurationCmd.MarkFlagRequired("channel-id")
	chatbot_updateMicrosoftTeamsChannelConfigurationCmd.MarkFlagRequired("chat-configuration-arn")
	chatbotCmd.AddCommand(chatbot_updateMicrosoftTeamsChannelConfigurationCmd)
}
