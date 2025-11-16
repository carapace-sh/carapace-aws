package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List all tags added to the specified Kinesis resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesis_listTagsForResourceCmd).Standalone()

		kinesis_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Kinesis resource for which to list tags.")
		kinesis_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	kinesisCmd.AddCommand(kinesis_listTagsForResourceCmd)
}
