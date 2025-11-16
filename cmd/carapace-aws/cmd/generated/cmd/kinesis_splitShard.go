package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_splitShardCmd = &cobra.Command{
	Use:   "split-shard",
	Short: "Splits a shard into two new shards in the Kinesis data stream, to increase the stream's capacity to ingest and transport data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_splitShardCmd).Standalone()

	kinesis_splitShardCmd.Flags().String("new-starting-hash-key", "", "A hash key value for the starting hash key of one of the child shards created by the split.")
	kinesis_splitShardCmd.Flags().String("shard-to-split", "", "The shard ID of the shard to split.")
	kinesis_splitShardCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
	kinesis_splitShardCmd.Flags().String("stream-name", "", "The name of the stream for the shard split.")
	kinesis_splitShardCmd.MarkFlagRequired("new-starting-hash-key")
	kinesis_splitShardCmd.MarkFlagRequired("shard-to-split")
	kinesisCmd.AddCommand(kinesis_splitShardCmd)
}
