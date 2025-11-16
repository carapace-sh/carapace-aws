package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Detaches a key-value pair from a resource, as identified by its Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_untagResourceCmd).Standalone()

	chatbot_untagResourceCmd.Flags().String("resource-arn", "", "The value of the resource that will have the tag removed.")
	chatbot_untagResourceCmd.Flags().String("tag-keys", "", "TagKeys are key-value pairs assigned to ARNs that can be used to group and search for resources by type.")
	chatbot_untagResourceCmd.MarkFlagRequired("resource-arn")
	chatbot_untagResourceCmd.MarkFlagRequired("tag-keys")
	chatbotCmd.AddCommand(chatbot_untagResourceCmd)
}
