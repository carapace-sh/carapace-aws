package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createDbclusterSnapshotCmd = &cobra.Command{
	Use:   "create-dbcluster-snapshot",
	Short: "Creates a snapshot of a DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createDbclusterSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_createDbclusterSnapshotCmd).Standalone()

		rds_createDbclusterSnapshotCmd.Flags().String("dbcluster-identifier", "", "The identifier of the DB cluster to create a snapshot for.")
		rds_createDbclusterSnapshotCmd.Flags().String("dbcluster-snapshot-identifier", "", "The identifier of the DB cluster snapshot.")
		rds_createDbclusterSnapshotCmd.Flags().String("tags", "", "The tags to be assigned to the DB cluster snapshot.")
		rds_createDbclusterSnapshotCmd.MarkFlagRequired("dbcluster-identifier")
		rds_createDbclusterSnapshotCmd.MarkFlagRequired("dbcluster-snapshot-identifier")
	})
	rdsCmd.AddCommand(rds_createDbclusterSnapshotCmd)
}
