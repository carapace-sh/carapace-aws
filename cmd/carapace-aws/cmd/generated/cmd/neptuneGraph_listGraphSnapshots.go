package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_listGraphSnapshotsCmd = &cobra.Command{
	Use:   "list-graph-snapshots",
	Short: "Lists available snapshots of a specified Neptune Analytics graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_listGraphSnapshotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptuneGraph_listGraphSnapshotsCmd).Standalone()

		neptuneGraph_listGraphSnapshotsCmd.Flags().String("graph-identifier", "", "The unique identifier of the Neptune Analytics graph.")
		neptuneGraph_listGraphSnapshotsCmd.Flags().String("max-results", "", "The total number of records to return in the command's output.")
		neptuneGraph_listGraphSnapshotsCmd.Flags().String("next-token", "", "Pagination token used to paginate output.")
	})
	neptuneGraphCmd.AddCommand(neptuneGraph_listGraphSnapshotsCmd)
}
