package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_createGraphSnapshotCmd = &cobra.Command{
	Use:   "create-graph-snapshot",
	Short: "Creates a snapshot of the specific graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_createGraphSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptuneGraph_createGraphSnapshotCmd).Standalone()

		neptuneGraph_createGraphSnapshotCmd.Flags().String("graph-identifier", "", "The unique identifier of the Neptune Analytics graph.")
		neptuneGraph_createGraphSnapshotCmd.Flags().String("snapshot-name", "", "The snapshot name.")
		neptuneGraph_createGraphSnapshotCmd.Flags().String("tags", "", "Adds metadata tags to the new graph.")
		neptuneGraph_createGraphSnapshotCmd.MarkFlagRequired("graph-identifier")
		neptuneGraph_createGraphSnapshotCmd.MarkFlagRequired("snapshot-name")
	})
	neptuneGraphCmd.AddCommand(neptuneGraph_createGraphSnapshotCmd)
}
