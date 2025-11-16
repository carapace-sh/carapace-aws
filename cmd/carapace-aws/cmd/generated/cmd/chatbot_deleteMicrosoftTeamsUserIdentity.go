package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_deleteMicrosoftTeamsUserIdentityCmd = &cobra.Command{
	Use:   "delete-microsoft-teams-user-identity",
	Short: "Identifes a user level permission for a channel configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_deleteMicrosoftTeamsUserIdentityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chatbot_deleteMicrosoftTeamsUserIdentityCmd).Standalone()

		chatbot_deleteMicrosoftTeamsUserIdentityCmd.Flags().String("chat-configuration-arn", "", "The ARN of the MicrosoftTeamsChannelConfiguration associated with the user identity to delete.")
		chatbot_deleteMicrosoftTeamsUserIdentityCmd.Flags().String("user-id", "", "The Microsoft Teams user ID.")
		chatbot_deleteMicrosoftTeamsUserIdentityCmd.MarkFlagRequired("chat-configuration-arn")
		chatbot_deleteMicrosoftTeamsUserIdentityCmd.MarkFlagRequired("user-id")
	})
	chatbotCmd.AddCommand(chatbot_deleteMicrosoftTeamsUserIdentityCmd)
}
