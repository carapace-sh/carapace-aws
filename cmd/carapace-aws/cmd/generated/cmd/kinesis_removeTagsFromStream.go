package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_removeTagsFromStreamCmd = &cobra.Command{
	Use:   "remove-tags-from-stream",
	Short: "Removes tags from the specified Kinesis data stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_removeTagsFromStreamCmd).Standalone()

	kinesis_removeTagsFromStreamCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
	kinesis_removeTagsFromStreamCmd.Flags().String("stream-name", "", "The name of the stream.")
	kinesis_removeTagsFromStreamCmd.Flags().String("tag-keys", "", "A list of tag keys.")
	kinesis_removeTagsFromStreamCmd.MarkFlagRequired("tag-keys")
	kinesisCmd.AddCommand(kinesis_removeTagsFromStreamCmd)
}
