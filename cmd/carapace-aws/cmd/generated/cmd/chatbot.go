package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbotCmd = &cobra.Command{
	Use:   "chatbot",
	Short: "The *AWS Chatbot API Reference* provides descriptions, API request parameters, and the XML response for each of the AWS Chatbot API actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chatbotCmd).Standalone()

	})
	rootCmd.AddCommand(chatbotCmd)
}
