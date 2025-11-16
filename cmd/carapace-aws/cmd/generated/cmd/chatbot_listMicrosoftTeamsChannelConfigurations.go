package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_listMicrosoftTeamsChannelConfigurationsCmd = &cobra.Command{
	Use:   "list-microsoft-teams-channel-configurations",
	Short: "Lists all AWS Chatbot Microsoft Teams channel configurations in an AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_listMicrosoftTeamsChannelConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chatbot_listMicrosoftTeamsChannelConfigurationsCmd).Standalone()

		chatbot_listMicrosoftTeamsChannelConfigurationsCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		chatbot_listMicrosoftTeamsChannelConfigurationsCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
		chatbot_listMicrosoftTeamsChannelConfigurationsCmd.Flags().String("team-id", "", "The ID of the Microsoft Teams authorized with AWS Chatbot.")
	})
	chatbotCmd.AddCommand(chatbot_listMicrosoftTeamsChannelConfigurationsCmd)
}
