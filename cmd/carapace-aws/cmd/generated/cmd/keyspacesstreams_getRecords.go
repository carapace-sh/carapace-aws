package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspacesstreams_getRecordsCmd = &cobra.Command{
	Use:   "get-records",
	Short: "Retrieves data records from a specified shard in an Amazon Keyspaces data stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspacesstreams_getRecordsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(keyspacesstreams_getRecordsCmd).Standalone()

		keyspacesstreams_getRecordsCmd.Flags().String("max-results", "", "The maximum number of records to return in a single `GetRecords` request.")
		keyspacesstreams_getRecordsCmd.Flags().String("shard-iterator", "", "The unique identifier of the shard iterator.")
		keyspacesstreams_getRecordsCmd.MarkFlagRequired("shard-iterator")
	})
	keyspacesstreamsCmd.AddCommand(keyspacesstreams_getRecordsCmd)
}
