package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteClusterSnapshotCmd = &cobra.Command{
	Use:   "delete-cluster-snapshot",
	Short: "Deletes the specified manual snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteClusterSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_deleteClusterSnapshotCmd).Standalone()

		redshift_deleteClusterSnapshotCmd.Flags().String("snapshot-cluster-identifier", "", "The unique identifier of the cluster the snapshot was created from.")
		redshift_deleteClusterSnapshotCmd.Flags().String("snapshot-identifier", "", "The unique identifier of the manual snapshot to be deleted.")
		redshift_deleteClusterSnapshotCmd.MarkFlagRequired("snapshot-identifier")
	})
	redshiftCmd.AddCommand(redshift_deleteClusterSnapshotCmd)
}
