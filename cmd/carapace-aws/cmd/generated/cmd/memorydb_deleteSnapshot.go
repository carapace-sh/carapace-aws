package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_deleteSnapshotCmd = &cobra.Command{
	Use:   "delete-snapshot",
	Short: "Deletes an existing snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_deleteSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_deleteSnapshotCmd).Standalone()

		memorydb_deleteSnapshotCmd.Flags().String("snapshot-name", "", "The name of the snapshot to delete.")
		memorydb_deleteSnapshotCmd.MarkFlagRequired("snapshot-name")
	})
	memorydbCmd.AddCommand(memorydb_deleteSnapshotCmd)
}
