package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_deleteGraphSnapshotCmd = &cobra.Command{
	Use:   "delete-graph-snapshot",
	Short: "Deletes the specifed graph snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_deleteGraphSnapshotCmd).Standalone()

	neptuneGraph_deleteGraphSnapshotCmd.Flags().String("snapshot-identifier", "", "ID of the graph snapshot to be deleted.")
	neptuneGraph_deleteGraphSnapshotCmd.MarkFlagRequired("snapshot-identifier")
	neptuneGraphCmd.AddCommand(neptuneGraph_deleteGraphSnapshotCmd)
}
