package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_deleteMicrosoftTeamsConfiguredTeamCmd = &cobra.Command{
	Use:   "delete-microsoft-teams-configured-team",
	Short: "Deletes the Microsoft Teams team authorization allowing for channels to be configured in that Microsoft Teams team.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_deleteMicrosoftTeamsConfiguredTeamCmd).Standalone()

	chatbot_deleteMicrosoftTeamsConfiguredTeamCmd.Flags().String("team-id", "", "The ID of the Microsoft Teams team authorized with AWS Chatbot.")
	chatbot_deleteMicrosoftTeamsConfiguredTeamCmd.MarkFlagRequired("team-id")
	chatbotCmd.AddCommand(chatbot_deleteMicrosoftTeamsConfiguredTeamCmd)
}
