package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspacesstreams_getStreamCmd = &cobra.Command{
	Use:   "get-stream",
	Short: "Returns detailed information about a specific data capture stream for an Amazon Keyspaces table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspacesstreams_getStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(keyspacesstreams_getStreamCmd).Standalone()

		keyspacesstreams_getStreamCmd.Flags().String("max-results", "", "The maximum number of shard objects to return in a single `GetStream` request.")
		keyspacesstreams_getStreamCmd.Flags().String("next-token", "", "An optional pagination token provided by a previous `GetStream` operation.")
		keyspacesstreams_getStreamCmd.Flags().String("shard-filter", "", "Optional filter criteria to apply when retrieving shards.")
		keyspacesstreams_getStreamCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream for which detailed information is requested.")
		keyspacesstreams_getStreamCmd.MarkFlagRequired("stream-arn")
	})
	keyspacesstreamsCmd.AddCommand(keyspacesstreams_getStreamCmd)
}
