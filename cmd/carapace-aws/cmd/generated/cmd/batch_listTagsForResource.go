package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags for an Batch resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_listTagsForResourceCmd).Standalone()

	batch_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that identifies the resource that tags are listed for.")
	batch_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	batchCmd.AddCommand(batch_listTagsForResourceCmd)
}
