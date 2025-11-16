package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_addTagsToStreamCmd = &cobra.Command{
	Use:   "add-tags-to-stream",
	Short: "Adds or updates tags for the specified Kinesis data stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_addTagsToStreamCmd).Standalone()

	kinesis_addTagsToStreamCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
	kinesis_addTagsToStreamCmd.Flags().String("stream-name", "", "The name of the stream.")
	kinesis_addTagsToStreamCmd.Flags().String("tags", "", "A set of up to 50 key-value pairs to use to create the tags.")
	kinesis_addTagsToStreamCmd.MarkFlagRequired("tags")
	kinesisCmd.AddCommand(kinesis_addTagsToStreamCmd)
}
