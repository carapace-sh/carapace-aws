package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspacesstreams_listStreamsCmd = &cobra.Command{
	Use:   "list-streams",
	Short: "Returns a list of all data capture streams associated with your Amazon Keyspaces account or for a specific keyspace or table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspacesstreams_listStreamsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(keyspacesstreams_listStreamsCmd).Standalone()

		keyspacesstreams_listStreamsCmd.Flags().String("keyspace-name", "", "The name of the keyspace for which to list streams.")
		keyspacesstreams_listStreamsCmd.Flags().String("max-results", "", "The maximum number of streams to return in a single `ListStreams` request.")
		keyspacesstreams_listStreamsCmd.Flags().String("next-token", "", "An optional pagination token provided by a previous `ListStreams` operation.")
		keyspacesstreams_listStreamsCmd.Flags().String("table-name", "", "The name of the table for which to list streams.")
	})
	keyspacesstreamsCmd.AddCommand(keyspacesstreams_listStreamsCmd)
}
