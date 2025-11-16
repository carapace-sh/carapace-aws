package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_listTagsForStreamCmd = &cobra.Command{
	Use:   "list-tags-for-stream",
	Short: "Lists the tags for the specified Kinesis data stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_listTagsForStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesis_listTagsForStreamCmd).Standalone()

		kinesis_listTagsForStreamCmd.Flags().String("exclusive-start-tag-key", "", "The key to use as the starting point for the list of tags.")
		kinesis_listTagsForStreamCmd.Flags().String("limit", "", "The number of tags to return.")
		kinesis_listTagsForStreamCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
		kinesis_listTagsForStreamCmd.Flags().String("stream-name", "", "The name of the stream.")
	})
	kinesisCmd.AddCommand(kinesis_listTagsForStreamCmd)
}
