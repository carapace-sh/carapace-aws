package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_listShardsCmd = &cobra.Command{
	Use:   "list-shards",
	Short: "Lists the shards in a stream and provides information about each shard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_listShardsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesis_listShardsCmd).Standalone()

		kinesis_listShardsCmd.Flags().String("exclusive-start-shard-id", "", "Specify this parameter to indicate that you want to list the shards starting with the shard whose ID immediately follows `ExclusiveStartShardId`.")
		kinesis_listShardsCmd.Flags().String("max-results", "", "The maximum number of shards to return in a single call to `ListShards`.")
		kinesis_listShardsCmd.Flags().String("next-token", "", "When the number of shards in the data stream is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of shards in the data stream, the response includes a pagination token named `NextToken`.")
		kinesis_listShardsCmd.Flags().String("shard-filter", "", "Enables you to filter out the response of the `ListShards` API.")
		kinesis_listShardsCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
		kinesis_listShardsCmd.Flags().String("stream-creation-timestamp", "", "Specify this input parameter to distinguish data streams that have the same name.")
		kinesis_listShardsCmd.Flags().String("stream-name", "", "The name of the data stream whose shards you want to list.")
	})
	kinesisCmd.AddCommand(kinesis_listShardsCmd)
}
