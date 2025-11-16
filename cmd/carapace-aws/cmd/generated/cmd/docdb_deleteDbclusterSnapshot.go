package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_deleteDbclusterSnapshotCmd = &cobra.Command{
	Use:   "delete-dbcluster-snapshot",
	Short: "Deletes a cluster snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_deleteDbclusterSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_deleteDbclusterSnapshotCmd).Standalone()

		docdb_deleteDbclusterSnapshotCmd.Flags().String("dbcluster-snapshot-identifier", "", "The identifier of the cluster snapshot to delete.")
		docdb_deleteDbclusterSnapshotCmd.MarkFlagRequired("dbcluster-snapshot-identifier")
	})
	docdbCmd.AddCommand(docdb_deleteDbclusterSnapshotCmd)
}
