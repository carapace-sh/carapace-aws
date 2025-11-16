package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove a tag from a Amazon SWF domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_untagResourceCmd).Standalone()

	swf_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the Amazon SWF domain.")
	swf_untagResourceCmd.Flags().String("tag-keys", "", "The list of tags to remove from the Amazon SWF domain.")
	swf_untagResourceCmd.MarkFlagRequired("resource-arn")
	swf_untagResourceCmd.MarkFlagRequired("tag-keys")
	swfCmd.AddCommand(swf_untagResourceCmd)
}
