package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_deleteChimeWebhookConfigurationCmd = &cobra.Command{
	Use:   "delete-chime-webhook-configuration",
	Short: "Deletes a Amazon Chime webhook configuration for AWS Chatbot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_deleteChimeWebhookConfigurationCmd).Standalone()

	chatbot_deleteChimeWebhookConfigurationCmd.Flags().String("chat-configuration-arn", "", "The Amazon Resource Name (ARN) of the ChimeWebhookConfiguration to delete.")
	chatbot_deleteChimeWebhookConfigurationCmd.MarkFlagRequired("chat-configuration-arn")
	chatbotCmd.AddCommand(chatbot_deleteChimeWebhookConfigurationCmd)
}
