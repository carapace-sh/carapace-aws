package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_copyClusterSnapshotCmd = &cobra.Command{
	Use:   "copy-cluster-snapshot",
	Short: "Copies the specified automated cluster snapshot to a new manual cluster snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_copyClusterSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_copyClusterSnapshotCmd).Standalone()

		redshift_copyClusterSnapshotCmd.Flags().String("manual-snapshot-retention-period", "", "The number of days that a manual snapshot is retained.")
		redshift_copyClusterSnapshotCmd.Flags().String("source-snapshot-cluster-identifier", "", "The identifier of the cluster the source snapshot was created from.")
		redshift_copyClusterSnapshotCmd.Flags().String("source-snapshot-identifier", "", "The identifier for the source snapshot.")
		redshift_copyClusterSnapshotCmd.Flags().String("target-snapshot-identifier", "", "The identifier given to the new manual snapshot.")
		redshift_copyClusterSnapshotCmd.MarkFlagRequired("source-snapshot-identifier")
		redshift_copyClusterSnapshotCmd.MarkFlagRequired("target-snapshot-identifier")
	})
	redshiftCmd.AddCommand(redshift_copyClusterSnapshotCmd)
}
