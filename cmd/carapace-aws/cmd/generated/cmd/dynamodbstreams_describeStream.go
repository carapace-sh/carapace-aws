package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodbstreams_describeStreamCmd = &cobra.Command{
	Use:   "describe-stream",
	Short: "Returns information about a stream, including the current status of the stream, its Amazon Resource Name (ARN), the composition of its shards, and its corresponding DynamoDB table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodbstreams_describeStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodbstreams_describeStreamCmd).Standalone()

		dynamodbstreams_describeStreamCmd.Flags().String("exclusive-start-shard-id", "", "The shard ID of the first item that this operation will evaluate.")
		dynamodbstreams_describeStreamCmd.Flags().String("limit", "", "The maximum number of shard objects to return.")
		dynamodbstreams_describeStreamCmd.Flags().String("shard-filter", "", "This optional field contains the filter definition for the `DescribeStream` API.")
		dynamodbstreams_describeStreamCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) for the stream.")
		dynamodbstreams_describeStreamCmd.MarkFlagRequired("stream-arn")
	})
	dynamodbstreamsCmd.AddCommand(dynamodbstreams_describeStreamCmd)
}
