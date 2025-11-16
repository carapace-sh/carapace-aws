package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_getMicrosoftTeamsChannelConfigurationCmd = &cobra.Command{
	Use:   "get-microsoft-teams-channel-configuration",
	Short: "Returns a Microsoft Teams channel configuration in an AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_getMicrosoftTeamsChannelConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chatbot_getMicrosoftTeamsChannelConfigurationCmd).Standalone()

		chatbot_getMicrosoftTeamsChannelConfigurationCmd.Flags().String("chat-configuration-arn", "", "The Amazon Resource Name (ARN) of the MicrosoftTeamsChannelConfiguration to retrieve.")
		chatbot_getMicrosoftTeamsChannelConfigurationCmd.MarkFlagRequired("chat-configuration-arn")
	})
	chatbotCmd.AddCommand(chatbot_getMicrosoftTeamsChannelConfigurationCmd)
}
