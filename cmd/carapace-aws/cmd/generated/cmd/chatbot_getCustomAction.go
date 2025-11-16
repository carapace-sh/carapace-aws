package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_getCustomActionCmd = &cobra.Command{
	Use:   "get-custom-action",
	Short: "Returns a custom action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_getCustomActionCmd).Standalone()

	chatbot_getCustomActionCmd.Flags().String("custom-action-arn", "", "Returns the fully defined Amazon Resource Name (ARN) of the custom action.")
	chatbot_getCustomActionCmd.MarkFlagRequired("custom-action-arn")
	chatbotCmd.AddCommand(chatbot_getCustomActionCmd)
}
