package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_subscribeToShardCmd = &cobra.Command{
	Use:   "subscribe-to-shard",
	Short: "This operation establishes an HTTP/2 connection between the consumer you specify in the `ConsumerARN` parameter and the shard you specify in the `ShardId` parameter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_subscribeToShardCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesis_subscribeToShardCmd).Standalone()

		kinesis_subscribeToShardCmd.Flags().String("consumer-arn", "", "For this parameter, use the value you obtained when you called [RegisterStreamConsumer]().")
		kinesis_subscribeToShardCmd.Flags().String("shard-id", "", "The ID of the shard you want to subscribe to.")
		kinesis_subscribeToShardCmd.Flags().String("starting-position", "", "The starting position in the data stream from which to start streaming.")
		kinesis_subscribeToShardCmd.MarkFlagRequired("consumer-arn")
		kinesis_subscribeToShardCmd.MarkFlagRequired("shard-id")
		kinesis_subscribeToShardCmd.MarkFlagRequired("starting-position")
	})
	kinesisCmd.AddCommand(kinesis_subscribeToShardCmd)
}
