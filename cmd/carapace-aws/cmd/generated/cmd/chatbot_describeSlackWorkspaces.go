package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_describeSlackWorkspacesCmd = &cobra.Command{
	Use:   "describe-slack-workspaces",
	Short: "List all authorized Slack workspaces connected to the AWS Account onboarded with AWS Chatbot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_describeSlackWorkspacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chatbot_describeSlackWorkspacesCmd).Standalone()

		chatbot_describeSlackWorkspacesCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		chatbot_describeSlackWorkspacesCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
	})
	chatbotCmd.AddCommand(chatbot_describeSlackWorkspacesCmd)
}
