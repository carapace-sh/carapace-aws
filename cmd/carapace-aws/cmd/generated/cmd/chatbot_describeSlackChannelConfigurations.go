package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_describeSlackChannelConfigurationsCmd = &cobra.Command{
	Use:   "describe-slack-channel-configurations",
	Short: "Lists Slack channel configurations optionally filtered by ChatConfigurationArn",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_describeSlackChannelConfigurationsCmd).Standalone()

	chatbot_describeSlackChannelConfigurationsCmd.Flags().String("chat-configuration-arn", "", "An optional Amazon Resource Name (ARN) of a SlackChannelConfiguration to describe.")
	chatbot_describeSlackChannelConfigurationsCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
	chatbot_describeSlackChannelConfigurationsCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
	chatbotCmd.AddCommand(chatbot_describeSlackChannelConfigurationsCmd)
}
