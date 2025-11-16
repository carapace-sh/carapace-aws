package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_deleteCustomActionCmd = &cobra.Command{
	Use:   "delete-custom-action",
	Short: "Deletes a custom action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_deleteCustomActionCmd).Standalone()

	chatbot_deleteCustomActionCmd.Flags().String("custom-action-arn", "", "The fully defined ARN of the custom action.")
	chatbot_deleteCustomActionCmd.MarkFlagRequired("custom-action-arn")
	chatbotCmd.AddCommand(chatbot_deleteCustomActionCmd)
}
