package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_restoreGraphFromSnapshotCmd = &cobra.Command{
	Use:   "restore-graph-from-snapshot",
	Short: "Restores a graph from a snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_restoreGraphFromSnapshotCmd).Standalone()

	neptuneGraph_restoreGraphFromSnapshotCmd.Flags().Bool("deletion-protection", false, "A value that indicates whether the graph has deletion protection enabled.")
	neptuneGraph_restoreGraphFromSnapshotCmd.Flags().String("graph-name", "", "A name for the new Neptune Analytics graph to be created from the snapshot.")
	neptuneGraph_restoreGraphFromSnapshotCmd.Flags().Bool("no-deletion-protection", false, "A value that indicates whether the graph has deletion protection enabled.")
	neptuneGraph_restoreGraphFromSnapshotCmd.Flags().Bool("no-public-connectivity", false, "Specifies whether or not the graph can be reachable over the internet.")
	neptuneGraph_restoreGraphFromSnapshotCmd.Flags().String("provisioned-memory", "", "The provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use for the graph.")
	neptuneGraph_restoreGraphFromSnapshotCmd.Flags().Bool("public-connectivity", false, "Specifies whether or not the graph can be reachable over the internet.")
	neptuneGraph_restoreGraphFromSnapshotCmd.Flags().String("replica-count", "", "The number of replicas in other AZs.")
	neptuneGraph_restoreGraphFromSnapshotCmd.Flags().String("snapshot-identifier", "", "The ID of the snapshot in question.")
	neptuneGraph_restoreGraphFromSnapshotCmd.Flags().String("tags", "", "Adds metadata tags to the snapshot.")
	neptuneGraph_restoreGraphFromSnapshotCmd.MarkFlagRequired("graph-name")
	neptuneGraph_restoreGraphFromSnapshotCmd.Flag("no-deletion-protection").Hidden = true
	neptuneGraph_restoreGraphFromSnapshotCmd.Flag("no-public-connectivity").Hidden = true
	neptuneGraph_restoreGraphFromSnapshotCmd.MarkFlagRequired("snapshot-identifier")
	neptuneGraphCmd.AddCommand(neptuneGraph_restoreGraphFromSnapshotCmd)
}
