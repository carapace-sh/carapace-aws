package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_deleteDbclusterSnapshotCmd = &cobra.Command{
	Use:   "delete-dbcluster-snapshot",
	Short: "Deletes a DB cluster snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_deleteDbclusterSnapshotCmd).Standalone()

	neptune_deleteDbclusterSnapshotCmd.Flags().String("dbcluster-snapshot-identifier", "", "The identifier of the DB cluster snapshot to delete.")
	neptune_deleteDbclusterSnapshotCmd.MarkFlagRequired("dbcluster-snapshot-identifier")
	neptuneCmd.AddCommand(neptune_deleteDbclusterSnapshotCmd)
}
