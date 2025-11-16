package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodbstreams_getShardIteratorCmd = &cobra.Command{
	Use:   "get-shard-iterator",
	Short: "Returns a shard iterator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodbstreams_getShardIteratorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodbstreams_getShardIteratorCmd).Standalone()

		dynamodbstreams_getShardIteratorCmd.Flags().String("sequence-number", "", "The sequence number of a stream record in the shard from which to start reading.")
		dynamodbstreams_getShardIteratorCmd.Flags().String("shard-id", "", "The identifier of the shard.")
		dynamodbstreams_getShardIteratorCmd.Flags().String("shard-iterator-type", "", "Determines how the shard iterator is used to start reading stream records from the shard:")
		dynamodbstreams_getShardIteratorCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) for the stream.")
		dynamodbstreams_getShardIteratorCmd.MarkFlagRequired("shard-id")
		dynamodbstreams_getShardIteratorCmd.MarkFlagRequired("shard-iterator-type")
		dynamodbstreams_getShardIteratorCmd.MarkFlagRequired("stream-arn")
	})
	dynamodbstreamsCmd.AddCommand(dynamodbstreams_getShardIteratorCmd)
}
