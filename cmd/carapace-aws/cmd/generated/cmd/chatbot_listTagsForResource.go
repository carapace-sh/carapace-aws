package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all of the tags associated with the Amazon Resource Name (ARN) that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_listTagsForResourceCmd).Standalone()

	chatbot_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to list tags for.")
	chatbot_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	chatbotCmd.AddCommand(chatbot_listTagsForResourceCmd)
}
