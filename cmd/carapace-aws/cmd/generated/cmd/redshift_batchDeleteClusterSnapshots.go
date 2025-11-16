package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_batchDeleteClusterSnapshotsCmd = &cobra.Command{
	Use:   "batch-delete-cluster-snapshots",
	Short: "Deletes a set of cluster snapshots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_batchDeleteClusterSnapshotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_batchDeleteClusterSnapshotsCmd).Standalone()

		redshift_batchDeleteClusterSnapshotsCmd.Flags().String("identifiers", "", "A list of identifiers for the snapshots that you want to delete.")
		redshift_batchDeleteClusterSnapshotsCmd.MarkFlagRequired("identifiers")
	})
	redshiftCmd.AddCommand(redshift_batchDeleteClusterSnapshotsCmd)
}
