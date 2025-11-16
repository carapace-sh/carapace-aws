package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Attaches a key-value pair to a resource, as identified by its Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_tagResourceCmd).Standalone()

	chatbot_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the configuration.")
	chatbot_tagResourceCmd.Flags().String("tags", "", "A list of tags to apply to the configuration.")
	chatbot_tagResourceCmd.MarkFlagRequired("resource-arn")
	chatbot_tagResourceCmd.MarkFlagRequired("tags")
	chatbotCmd.AddCommand(chatbot_tagResourceCmd)
}
