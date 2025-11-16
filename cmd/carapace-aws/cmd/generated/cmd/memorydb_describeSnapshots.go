package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_describeSnapshotsCmd = &cobra.Command{
	Use:   "describe-snapshots",
	Short: "Returns information about cluster snapshots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_describeSnapshotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_describeSnapshotsCmd).Standalone()

		memorydb_describeSnapshotsCmd.Flags().String("cluster-name", "", "A user-supplied cluster identifier.")
		memorydb_describeSnapshotsCmd.Flags().String("max-results", "", "The maximum number of records to include in the response.")
		memorydb_describeSnapshotsCmd.Flags().String("next-token", "", "An optional argument to pass in case the total number of records exceeds the value of MaxResults.")
		memorydb_describeSnapshotsCmd.Flags().String("show-detail", "", "A Boolean value which if true, the shard configuration is included in the snapshot description.")
		memorydb_describeSnapshotsCmd.Flags().String("snapshot-name", "", "A user-supplied name of the snapshot.")
		memorydb_describeSnapshotsCmd.Flags().String("source", "", "If set to system, the output shows snapshots that were automatically created by MemoryDB.")
	})
	memorydbCmd.AddCommand(memorydb_describeSnapshotsCmd)
}
