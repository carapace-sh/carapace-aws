package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_createClusterSnapshotCmd = &cobra.Command{
	Use:   "create-cluster-snapshot",
	Short: "Creates a manual snapshot of the specified cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_createClusterSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_createClusterSnapshotCmd).Standalone()

		redshift_createClusterSnapshotCmd.Flags().String("cluster-identifier", "", "The cluster identifier for which you want a snapshot.")
		redshift_createClusterSnapshotCmd.Flags().String("manual-snapshot-retention-period", "", "The number of days that a manual snapshot is retained.")
		redshift_createClusterSnapshotCmd.Flags().String("snapshot-identifier", "", "A unique identifier for the snapshot that you are requesting.")
		redshift_createClusterSnapshotCmd.Flags().String("tags", "", "A list of tag instances.")
		redshift_createClusterSnapshotCmd.MarkFlagRequired("cluster-identifier")
		redshift_createClusterSnapshotCmd.MarkFlagRequired("snapshot-identifier")
	})
	redshiftCmd.AddCommand(redshift_createClusterSnapshotCmd)
}
