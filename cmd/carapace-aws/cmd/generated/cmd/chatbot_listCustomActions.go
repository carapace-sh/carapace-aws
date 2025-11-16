package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_listCustomActionsCmd = &cobra.Command{
	Use:   "list-custom-actions",
	Short: "Lists custom actions defined in this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_listCustomActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chatbot_listCustomActionsCmd).Standalone()

		chatbot_listCustomActionsCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		chatbot_listCustomActionsCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
	})
	chatbotCmd.AddCommand(chatbot_listCustomActionsCmd)
}
