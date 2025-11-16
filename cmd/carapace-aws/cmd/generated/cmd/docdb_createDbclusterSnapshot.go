package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_createDbclusterSnapshotCmd = &cobra.Command{
	Use:   "create-dbcluster-snapshot",
	Short: "Creates a snapshot of a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_createDbclusterSnapshotCmd).Standalone()

	docdb_createDbclusterSnapshotCmd.Flags().String("dbcluster-identifier", "", "The identifier of the cluster to create a snapshot for.")
	docdb_createDbclusterSnapshotCmd.Flags().String("dbcluster-snapshot-identifier", "", "The identifier of the cluster snapshot.")
	docdb_createDbclusterSnapshotCmd.Flags().String("tags", "", "The tags to be assigned to the cluster snapshot.")
	docdb_createDbclusterSnapshotCmd.MarkFlagRequired("dbcluster-identifier")
	docdb_createDbclusterSnapshotCmd.MarkFlagRequired("dbcluster-snapshot-identifier")
	docdbCmd.AddCommand(docdb_createDbclusterSnapshotCmd)
}
