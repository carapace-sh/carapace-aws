package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_listMicrosoftTeamsUserIdentitiesCmd = &cobra.Command{
	Use:   "list-microsoft-teams-user-identities",
	Short: "A list all Microsoft Teams user identities with a mapped role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_listMicrosoftTeamsUserIdentitiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chatbot_listMicrosoftTeamsUserIdentitiesCmd).Standalone()

		chatbot_listMicrosoftTeamsUserIdentitiesCmd.Flags().String("chat-configuration-arn", "", "The Amazon Resource Name (ARN) of the MicrosoftTeamsChannelConfiguration associated with the user identities to list.")
		chatbot_listMicrosoftTeamsUserIdentitiesCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		chatbot_listMicrosoftTeamsUserIdentitiesCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
	})
	chatbotCmd.AddCommand(chatbot_listMicrosoftTeamsUserIdentitiesCmd)
}
