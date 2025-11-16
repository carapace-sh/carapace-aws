package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_listMicrosoftTeamsConfiguredTeamsCmd = &cobra.Command{
	Use:   "list-microsoft-teams-configured-teams",
	Short: "Lists all authorized Microsoft Teams for an AWS Account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_listMicrosoftTeamsConfiguredTeamsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chatbot_listMicrosoftTeamsConfiguredTeamsCmd).Standalone()

		chatbot_listMicrosoftTeamsConfiguredTeamsCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		chatbot_listMicrosoftTeamsConfiguredTeamsCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
	})
	chatbotCmd.AddCommand(chatbot_listMicrosoftTeamsConfiguredTeamsCmd)
}
