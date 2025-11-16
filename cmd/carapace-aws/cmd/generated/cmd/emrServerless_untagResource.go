package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrServerless_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrServerless_untagResourceCmd).Standalone()

	emrServerless_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that identifies the resource to list the tags for.")
	emrServerless_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to be removed.")
	emrServerless_untagResourceCmd.MarkFlagRequired("resource-arn")
	emrServerless_untagResourceCmd.MarkFlagRequired("tag-keys")
	emrServerlessCmd.AddCommand(emrServerless_untagResourceCmd)
}
