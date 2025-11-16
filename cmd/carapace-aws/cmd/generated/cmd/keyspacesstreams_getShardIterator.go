package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspacesstreams_getShardIteratorCmd = &cobra.Command{
	Use:   "get-shard-iterator",
	Short: "Returns a shard iterator that serves as a bookmark for reading data from a specific position in an Amazon Keyspaces data stream's shard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspacesstreams_getShardIteratorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(keyspacesstreams_getShardIteratorCmd).Standalone()

		keyspacesstreams_getShardIteratorCmd.Flags().String("sequence-number", "", "The sequence number of the data record in the shard from which to start reading.")
		keyspacesstreams_getShardIteratorCmd.Flags().String("shard-id", "", "The identifier of the shard within the stream.")
		keyspacesstreams_getShardIteratorCmd.Flags().String("shard-iterator-type", "", "Determines how the shard iterator is positioned.")
		keyspacesstreams_getShardIteratorCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream for which to get the shard iterator.")
		keyspacesstreams_getShardIteratorCmd.MarkFlagRequired("shard-id")
		keyspacesstreams_getShardIteratorCmd.MarkFlagRequired("shard-iterator-type")
		keyspacesstreams_getShardIteratorCmd.MarkFlagRequired("stream-arn")
	})
	keyspacesstreamsCmd.AddCommand(keyspacesstreams_getShardIteratorCmd)
}
