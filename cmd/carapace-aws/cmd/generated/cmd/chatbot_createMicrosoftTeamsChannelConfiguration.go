package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_createMicrosoftTeamsChannelConfigurationCmd = &cobra.Command{
	Use:   "create-microsoft-teams-channel-configuration",
	Short: "Creates an AWS Chatbot configuration for Microsoft Teams.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_createMicrosoftTeamsChannelConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chatbot_createMicrosoftTeamsChannelConfigurationCmd).Standalone()

		chatbot_createMicrosoftTeamsChannelConfigurationCmd.Flags().String("channel-id", "", "The ID of the Microsoft Teams channel.")
		chatbot_createMicrosoftTeamsChannelConfigurationCmd.Flags().String("channel-name", "", "The name of the Microsoft Teams channel.")
		chatbot_createMicrosoftTeamsChannelConfigurationCmd.Flags().String("configuration-name", "", "The name of the configuration.")
		chatbot_createMicrosoftTeamsChannelConfigurationCmd.Flags().String("guardrail-policy-arns", "", "The list of IAM policy ARNs that are applied as channel guardrails.")
		chatbot_createMicrosoftTeamsChannelConfigurationCmd.Flags().String("iam-role-arn", "", "A user-defined role that AWS Chatbot assumes.")
		chatbot_createMicrosoftTeamsChannelConfigurationCmd.Flags().String("logging-level", "", "Logging levels include `ERROR`, `INFO`, or `NONE`.")
		chatbot_createMicrosoftTeamsChannelConfigurationCmd.Flags().String("sns-topic-arns", "", "The Amazon Resource Names (ARNs) of the SNS topics that deliver notifications to AWS Chatbot.")
		chatbot_createMicrosoftTeamsChannelConfigurationCmd.Flags().String("tags", "", "A map of tags assigned to a resource.")
		chatbot_createMicrosoftTeamsChannelConfigurationCmd.Flags().String("team-id", "", "The ID of the Microsoft Teams authorized with AWS Chatbot.")
		chatbot_createMicrosoftTeamsChannelConfigurationCmd.Flags().String("team-name", "", "The name of the Microsoft Teams Team.")
		chatbot_createMicrosoftTeamsChannelConfigurationCmd.Flags().String("tenant-id", "", "The ID of the Microsoft Teams tenant.")
		chatbot_createMicrosoftTeamsChannelConfigurationCmd.Flags().String("user-authorization-required", "", "Enables use of a user role requirement in your chat configuration.")
		chatbot_createMicrosoftTeamsChannelConfigurationCmd.MarkFlagRequired("channel-id")
		chatbot_createMicrosoftTeamsChannelConfigurationCmd.MarkFlagRequired("configuration-name")
		chatbot_createMicrosoftTeamsChannelConfigurationCmd.MarkFlagRequired("iam-role-arn")
		chatbot_createMicrosoftTeamsChannelConfigurationCmd.MarkFlagRequired("team-id")
		chatbot_createMicrosoftTeamsChannelConfigurationCmd.MarkFlagRequired("tenant-id")
	})
	chatbotCmd.AddCommand(chatbot_createMicrosoftTeamsChannelConfigurationCmd)
}
