package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_createDbclusterSnapshotCmd = &cobra.Command{
	Use:   "create-dbcluster-snapshot",
	Short: "Creates a snapshot of a DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_createDbclusterSnapshotCmd).Standalone()

	neptune_createDbclusterSnapshotCmd.Flags().String("dbcluster-identifier", "", "The identifier of the DB cluster to create a snapshot for.")
	neptune_createDbclusterSnapshotCmd.Flags().String("dbcluster-snapshot-identifier", "", "The identifier of the DB cluster snapshot.")
	neptune_createDbclusterSnapshotCmd.Flags().String("tags", "", "The tags to be assigned to the DB cluster snapshot.")
	neptune_createDbclusterSnapshotCmd.MarkFlagRequired("dbcluster-identifier")
	neptune_createDbclusterSnapshotCmd.MarkFlagRequired("dbcluster-snapshot-identifier")
	neptuneCmd.AddCommand(neptune_createDbclusterSnapshotCmd)
}
