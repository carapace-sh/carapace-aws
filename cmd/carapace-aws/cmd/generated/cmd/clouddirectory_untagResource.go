package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "An API operation for removing tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_untagResourceCmd).Standalone()

	clouddirectory_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	clouddirectory_untagResourceCmd.Flags().String("tag-keys", "", "Keys of the tag that need to be removed from the resource.")
	clouddirectory_untagResourceCmd.MarkFlagRequired("resource-arn")
	clouddirectory_untagResourceCmd.MarkFlagRequired("tag-keys")
	clouddirectoryCmd.AddCommand(clouddirectory_untagResourceCmd)
}
