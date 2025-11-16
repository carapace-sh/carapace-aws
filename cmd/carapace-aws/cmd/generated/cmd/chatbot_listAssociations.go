package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_listAssociationsCmd = &cobra.Command{
	Use:   "list-associations",
	Short: "Lists resources associated with a channel configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_listAssociationsCmd).Standalone()

	chatbot_listAssociationsCmd.Flags().String("chat-configuration", "", "The channel configuration to list associations for.")
	chatbot_listAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
	chatbot_listAssociationsCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
	chatbot_listAssociationsCmd.MarkFlagRequired("chat-configuration")
	chatbotCmd.AddCommand(chatbot_listAssociationsCmd)
}
