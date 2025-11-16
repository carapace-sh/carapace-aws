package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_getGraphSnapshotCmd = &cobra.Command{
	Use:   "get-graph-snapshot",
	Short: "Retrieves a specified graph snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_getGraphSnapshotCmd).Standalone()

	neptuneGraph_getGraphSnapshotCmd.Flags().String("snapshot-identifier", "", "The ID of the snapshot to retrieve.")
	neptuneGraph_getGraphSnapshotCmd.MarkFlagRequired("snapshot-identifier")
	neptuneGraphCmd.AddCommand(neptuneGraph_getGraphSnapshotCmd)
}
