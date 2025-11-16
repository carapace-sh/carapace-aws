package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tag key and value pairs from an Amazon Web Services Resource Explorer view or index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceExplorer2_untagResourceCmd).Standalone()

		resourceExplorer2_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the view or index that you want to remove tags from.")
		resourceExplorer2_untagResourceCmd.Flags().String("tag-keys", "", "A list of the keys for the tags that you want to remove from the specified view or index.")
		resourceExplorer2_untagResourceCmd.MarkFlagRequired("resource-arn")
		resourceExplorer2_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_untagResourceCmd)
}
