package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_updateCustomActionCmd = &cobra.Command{
	Use:   "update-custom-action",
	Short: "Updates a custom action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_updateCustomActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chatbot_updateCustomActionCmd).Standalone()

		chatbot_updateCustomActionCmd.Flags().String("alias-name", "", "The name used to invoke this action in the chat channel.")
		chatbot_updateCustomActionCmd.Flags().String("attachments", "", "Defines when this custom action button should be attached to a notification.")
		chatbot_updateCustomActionCmd.Flags().String("custom-action-arn", "", "The fully defined Amazon Resource Name (ARN) of the custom action.")
		chatbot_updateCustomActionCmd.Flags().String("definition", "", "The definition of the command to run when invoked as an alias or as an action button.")
		chatbot_updateCustomActionCmd.MarkFlagRequired("custom-action-arn")
		chatbot_updateCustomActionCmd.MarkFlagRequired("definition")
	})
	chatbotCmd.AddCommand(chatbot_updateCustomActionCmd)
}
