package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_getShardIteratorCmd = &cobra.Command{
	Use:   "get-shard-iterator",
	Short: "Gets an Amazon Kinesis shard iterator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_getShardIteratorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesis_getShardIteratorCmd).Standalone()

		kinesis_getShardIteratorCmd.Flags().String("shard-id", "", "The shard ID of the Kinesis Data Streams shard to get the iterator for.")
		kinesis_getShardIteratorCmd.Flags().String("shard-iterator-type", "", "Determines how the shard iterator is used to start reading data records from the shard.")
		kinesis_getShardIteratorCmd.Flags().String("starting-sequence-number", "", "The sequence number of the data record in the shard from which to start reading.")
		kinesis_getShardIteratorCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
		kinesis_getShardIteratorCmd.Flags().String("stream-name", "", "The name of the Amazon Kinesis data stream.")
		kinesis_getShardIteratorCmd.Flags().String("timestamp", "", "The time stamp of the data record from which to start reading.")
		kinesis_getShardIteratorCmd.MarkFlagRequired("shard-id")
		kinesis_getShardIteratorCmd.MarkFlagRequired("shard-iterator-type")
	})
	kinesisCmd.AddCommand(kinesis_getShardIteratorCmd)
}
