package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove one or more tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_untagResourceCmd).Standalone()

	bedrock_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to untag.")
	bedrock_untagResourceCmd.Flags().String("tag-keys", "", "Tag keys of the tags to remove from the resource.")
	bedrock_untagResourceCmd.MarkFlagRequired("resource-arn")
	bedrock_untagResourceCmd.MarkFlagRequired("tag-keys")
	bedrockCmd.AddCommand(bedrock_untagResourceCmd)
}
