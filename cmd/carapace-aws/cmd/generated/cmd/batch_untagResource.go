package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes specified tags from an Batch resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_untagResourceCmd).Standalone()

	batch_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource from which to delete tags.")
	batch_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to be removed.")
	batch_untagResourceCmd.MarkFlagRequired("resource-arn")
	batch_untagResourceCmd.MarkFlagRequired("tag-keys")
	batchCmd.AddCommand(batch_untagResourceCmd)
}
