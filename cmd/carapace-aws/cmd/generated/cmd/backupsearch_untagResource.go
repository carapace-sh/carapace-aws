package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupsearch_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "This operation removes tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupsearch_untagResourceCmd).Standalone()

	backupsearch_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the resource where you want to remove tags.")
	backupsearch_untagResourceCmd.Flags().String("tag-keys", "", "This required parameter contains the tag keys you want to remove from the source.")
	backupsearch_untagResourceCmd.MarkFlagRequired("resource-arn")
	backupsearch_untagResourceCmd.MarkFlagRequired("tag-keys")
	backupsearchCmd.AddCommand(backupsearch_untagResourceCmd)
}
