package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_updateChimeWebhookConfigurationCmd = &cobra.Command{
	Use:   "update-chime-webhook-configuration",
	Short: "Updates a Amazon Chime webhook configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_updateChimeWebhookConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chatbot_updateChimeWebhookConfigurationCmd).Standalone()

		chatbot_updateChimeWebhookConfigurationCmd.Flags().String("chat-configuration-arn", "", "The Amazon Resource Name (ARN) of the ChimeWebhookConfiguration to update.")
		chatbot_updateChimeWebhookConfigurationCmd.Flags().String("iam-role-arn", "", "A user-defined role that AWS Chatbot assumes.")
		chatbot_updateChimeWebhookConfigurationCmd.Flags().String("logging-level", "", "Logging levels include `ERROR`, `INFO`, or `NONE`.")
		chatbot_updateChimeWebhookConfigurationCmd.Flags().String("sns-topic-arns", "", "The ARNs of the SNS topics that deliver notifications to AWS Chatbot.")
		chatbot_updateChimeWebhookConfigurationCmd.Flags().String("webhook-description", "", "A description of the webhook.")
		chatbot_updateChimeWebhookConfigurationCmd.Flags().String("webhook-url", "", "The URL for the Amazon Chime webhook.")
		chatbot_updateChimeWebhookConfigurationCmd.MarkFlagRequired("chat-configuration-arn")
	})
	chatbotCmd.AddCommand(chatbot_updateChimeWebhookConfigurationCmd)
}
