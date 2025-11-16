package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodbstreams_getRecordsCmd = &cobra.Command{
	Use:   "get-records",
	Short: "Retrieves the stream records from a given shard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodbstreams_getRecordsCmd).Standalone()

	dynamodbstreams_getRecordsCmd.Flags().String("limit", "", "The maximum number of records to return from the shard.")
	dynamodbstreams_getRecordsCmd.Flags().String("shard-iterator", "", "A shard iterator that was retrieved from a previous GetShardIterator operation.")
	dynamodbstreams_getRecordsCmd.MarkFlagRequired("shard-iterator")
	dynamodbstreamsCmd.AddCommand(dynamodbstreams_getRecordsCmd)
}
