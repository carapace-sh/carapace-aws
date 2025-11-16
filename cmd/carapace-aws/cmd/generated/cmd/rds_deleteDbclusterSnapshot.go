package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteDbclusterSnapshotCmd = &cobra.Command{
	Use:   "delete-dbcluster-snapshot",
	Short: "Deletes a DB cluster snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteDbclusterSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_deleteDbclusterSnapshotCmd).Standalone()

		rds_deleteDbclusterSnapshotCmd.Flags().String("dbcluster-snapshot-identifier", "", "The identifier of the DB cluster snapshot to delete.")
		rds_deleteDbclusterSnapshotCmd.MarkFlagRequired("dbcluster-snapshot-identifier")
	})
	rdsCmd.AddCommand(rds_deleteDbclusterSnapshotCmd)
}
