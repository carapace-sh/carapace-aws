package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_getAccountPreferencesCmd = &cobra.Command{
	Use:   "get-account-preferences",
	Short: "Returns AWS Chatbot account preferences.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_getAccountPreferencesCmd).Standalone()

	chatbotCmd.AddCommand(chatbot_getAccountPreferencesCmd)
}
