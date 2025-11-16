package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_createCustomActionCmd = &cobra.Command{
	Use:   "create-custom-action",
	Short: "Creates a custom action that can be invoked as an alias or as a button on a notification.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_createCustomActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chatbot_createCustomActionCmd).Standalone()

		chatbot_createCustomActionCmd.Flags().String("action-name", "", "The name of the custom action.")
		chatbot_createCustomActionCmd.Flags().String("alias-name", "", "The name used to invoke this action in a chat channel.")
		chatbot_createCustomActionCmd.Flags().String("attachments", "", "Defines when this custom action button should be attached to a notification.")
		chatbot_createCustomActionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		chatbot_createCustomActionCmd.Flags().String("definition", "", "The definition of the command to run when invoked as an alias or as an action button.")
		chatbot_createCustomActionCmd.Flags().String("tags", "", "A map of tags assigned to a resource.")
		chatbot_createCustomActionCmd.MarkFlagRequired("action-name")
		chatbot_createCustomActionCmd.MarkFlagRequired("definition")
	})
	chatbotCmd.AddCommand(chatbot_createCustomActionCmd)
}
