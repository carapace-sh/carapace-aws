package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_updateAccountPreferencesCmd = &cobra.Command{
	Use:   "update-account-preferences",
	Short: "Updates AWS Chatbot account preferences.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_updateAccountPreferencesCmd).Standalone()

	chatbot_updateAccountPreferencesCmd.Flags().String("training-data-collection-enabled", "", "Turns on training data collection.")
	chatbot_updateAccountPreferencesCmd.Flags().String("user-authorization-required", "", "Enables use of a user role requirement in your chat configuration.")
	chatbotCmd.AddCommand(chatbot_updateAccountPreferencesCmd)
}
