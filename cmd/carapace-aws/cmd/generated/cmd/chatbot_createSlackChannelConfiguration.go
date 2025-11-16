package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_createSlackChannelConfigurationCmd = &cobra.Command{
	Use:   "create-slack-channel-configuration",
	Short: "Creates an AWS Chatbot confugration for Slack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_createSlackChannelConfigurationCmd).Standalone()

	chatbot_createSlackChannelConfigurationCmd.Flags().String("configuration-name", "", "The name of the configuration.")
	chatbot_createSlackChannelConfigurationCmd.Flags().String("guardrail-policy-arns", "", "The list of IAM policy ARNs that are applied as channel guardrails.")
	chatbot_createSlackChannelConfigurationCmd.Flags().String("iam-role-arn", "", "A user-defined role that AWS Chatbot assumes.")
	chatbot_createSlackChannelConfigurationCmd.Flags().String("logging-level", "", "Logging levels include `ERROR`, `INFO`, or `NONE`.")
	chatbot_createSlackChannelConfigurationCmd.Flags().String("slack-channel-id", "", "The ID of the Slack channel.")
	chatbot_createSlackChannelConfigurationCmd.Flags().String("slack-channel-name", "", "The name of the Slack channel.")
	chatbot_createSlackChannelConfigurationCmd.Flags().String("slack-team-id", "", "The ID of the Slack workspace authorized with AWS Chatbot.")
	chatbot_createSlackChannelConfigurationCmd.Flags().String("sns-topic-arns", "", "The Amazon Resource Names (ARNs) of the SNS topics that deliver notifications to AWS Chatbot.")
	chatbot_createSlackChannelConfigurationCmd.Flags().String("tags", "", "A map of tags assigned to a resource.")
	chatbot_createSlackChannelConfigurationCmd.Flags().String("user-authorization-required", "", "Enables use of a user role requirement in your chat configuration.")
	chatbot_createSlackChannelConfigurationCmd.MarkFlagRequired("configuration-name")
	chatbot_createSlackChannelConfigurationCmd.MarkFlagRequired("iam-role-arn")
	chatbot_createSlackChannelConfigurationCmd.MarkFlagRequired("slack-channel-id")
	chatbot_createSlackChannelConfigurationCmd.MarkFlagRequired("slack-team-id")
	chatbotCmd.AddCommand(chatbot_createSlackChannelConfigurationCmd)
}
