package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_describeChimeWebhookConfigurationsCmd = &cobra.Command{
	Use:   "describe-chime-webhook-configurations",
	Short: "Lists Amazon Chime webhook configurations optionally filtered by ChatConfigurationArn",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_describeChimeWebhookConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chatbot_describeChimeWebhookConfigurationsCmd).Standalone()

		chatbot_describeChimeWebhookConfigurationsCmd.Flags().String("chat-configuration-arn", "", "An optional Amazon Resource Name (ARN) of a ChimeWebhookConfiguration to describe.")
		chatbot_describeChimeWebhookConfigurationsCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		chatbot_describeChimeWebhookConfigurationsCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
	})
	chatbotCmd.AddCommand(chatbot_describeChimeWebhookConfigurationsCmd)
}
