package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_deleteMicrosoftTeamsChannelConfigurationCmd = &cobra.Command{
	Use:   "delete-microsoft-teams-channel-configuration",
	Short: "Deletes a Microsoft Teams channel configuration for AWS Chatbot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_deleteMicrosoftTeamsChannelConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chatbot_deleteMicrosoftTeamsChannelConfigurationCmd).Standalone()

		chatbot_deleteMicrosoftTeamsChannelConfigurationCmd.Flags().String("chat-configuration-arn", "", "The Amazon Resource Name (ARN) of the MicrosoftTeamsChannelConfiguration associated with the user identity to delete.")
		chatbot_deleteMicrosoftTeamsChannelConfigurationCmd.MarkFlagRequired("chat-configuration-arn")
	})
	chatbotCmd.AddCommand(chatbot_deleteMicrosoftTeamsChannelConfigurationCmd)
}
