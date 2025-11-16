package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_updateShardCountCmd = &cobra.Command{
	Use:   "update-shard-count",
	Short: "Updates the shard count of the specified stream to the specified number of shards.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_updateShardCountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesis_updateShardCountCmd).Standalone()

		kinesis_updateShardCountCmd.Flags().String("scaling-type", "", "The scaling type.")
		kinesis_updateShardCountCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
		kinesis_updateShardCountCmd.Flags().String("stream-name", "", "The name of the stream.")
		kinesis_updateShardCountCmd.Flags().String("target-shard-count", "", "The new number of shards.")
		kinesis_updateShardCountCmd.MarkFlagRequired("scaling-type")
		kinesis_updateShardCountCmd.MarkFlagRequired("target-shard-count")
	})
	kinesisCmd.AddCommand(kinesis_updateShardCountCmd)
}
