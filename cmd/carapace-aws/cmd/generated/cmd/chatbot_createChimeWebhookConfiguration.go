package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_createChimeWebhookConfigurationCmd = &cobra.Command{
	Use:   "create-chime-webhook-configuration",
	Short: "Creates an AWS Chatbot configuration for Amazon Chime.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_createChimeWebhookConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chatbot_createChimeWebhookConfigurationCmd).Standalone()

		chatbot_createChimeWebhookConfigurationCmd.Flags().String("configuration-name", "", "The name of the configuration.")
		chatbot_createChimeWebhookConfigurationCmd.Flags().String("iam-role-arn", "", "A user-defined role that AWS Chatbot assumes.")
		chatbot_createChimeWebhookConfigurationCmd.Flags().String("logging-level", "", "Logging levels include `ERROR`, `INFO`, or `NONE`.")
		chatbot_createChimeWebhookConfigurationCmd.Flags().String("sns-topic-arns", "", "The Amazon Resource Names (ARNs) of the SNS topics that deliver notifications to AWS Chatbot.")
		chatbot_createChimeWebhookConfigurationCmd.Flags().String("tags", "", "A map of tags assigned to a resource.")
		chatbot_createChimeWebhookConfigurationCmd.Flags().String("webhook-description", "", "A description of the webhook.")
		chatbot_createChimeWebhookConfigurationCmd.Flags().String("webhook-url", "", "The URL for the Amazon Chime webhook.")
		chatbot_createChimeWebhookConfigurationCmd.MarkFlagRequired("configuration-name")
		chatbot_createChimeWebhookConfigurationCmd.MarkFlagRequired("iam-role-arn")
		chatbot_createChimeWebhookConfigurationCmd.MarkFlagRequired("sns-topic-arns")
		chatbot_createChimeWebhookConfigurationCmd.MarkFlagRequired("webhook-description")
		chatbot_createChimeWebhookConfigurationCmd.MarkFlagRequired("webhook-url")
	})
	chatbotCmd.AddCommand(chatbot_createChimeWebhookConfigurationCmd)
}
