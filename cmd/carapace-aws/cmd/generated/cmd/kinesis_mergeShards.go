package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_mergeShardsCmd = &cobra.Command{
	Use:   "merge-shards",
	Short: "Merges two adjacent shards in a Kinesis data stream and combines them into a single shard to reduce the stream's capacity to ingest and transport data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_mergeShardsCmd).Standalone()

	kinesis_mergeShardsCmd.Flags().String("adjacent-shard-to-merge", "", "The shard ID of the adjacent shard for the merge.")
	kinesis_mergeShardsCmd.Flags().String("shard-to-merge", "", "The shard ID of the shard to combine with the adjacent shard for the merge.")
	kinesis_mergeShardsCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
	kinesis_mergeShardsCmd.Flags().String("stream-name", "", "The name of the stream for the merge.")
	kinesis_mergeShardsCmd.MarkFlagRequired("adjacent-shard-to-merge")
	kinesis_mergeShardsCmd.MarkFlagRequired("shard-to-merge")
	kinesisCmd.AddCommand(kinesis_mergeShardsCmd)
}
