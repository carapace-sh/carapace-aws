package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_restoreFromSnapshotCmd = &cobra.Command{
	Use:   "restore-from-snapshot",
	Short: "Restores a directory using an existing directory snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_restoreFromSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_restoreFromSnapshotCmd).Standalone()

		ds_restoreFromSnapshotCmd.Flags().String("snapshot-id", "", "The identifier of the snapshot to restore from.")
		ds_restoreFromSnapshotCmd.MarkFlagRequired("snapshot-id")
	})
	dsCmd.AddCommand(ds_restoreFromSnapshotCmd)
}
