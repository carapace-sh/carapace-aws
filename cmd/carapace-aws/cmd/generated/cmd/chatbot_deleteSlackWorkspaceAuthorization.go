package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_deleteSlackWorkspaceAuthorizationCmd = &cobra.Command{
	Use:   "delete-slack-workspace-authorization",
	Short: "Deletes the Slack workspace authorization that allows channels to be configured in that workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_deleteSlackWorkspaceAuthorizationCmd).Standalone()

	chatbot_deleteSlackWorkspaceAuthorizationCmd.Flags().String("slack-team-id", "", "The ID of the Slack workspace authorized with AWS Chatbot.")
	chatbot_deleteSlackWorkspaceAuthorizationCmd.MarkFlagRequired("slack-team-id")
	chatbotCmd.AddCommand(chatbot_deleteSlackWorkspaceAuthorizationCmd)
}
