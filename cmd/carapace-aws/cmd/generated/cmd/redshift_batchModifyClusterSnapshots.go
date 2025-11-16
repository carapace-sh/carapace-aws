package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_batchModifyClusterSnapshotsCmd = &cobra.Command{
	Use:   "batch-modify-cluster-snapshots",
	Short: "Modifies the settings for a set of cluster snapshots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_batchModifyClusterSnapshotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_batchModifyClusterSnapshotsCmd).Standalone()

		redshift_batchModifyClusterSnapshotsCmd.Flags().Bool("force", false, "A boolean value indicating whether to override an exception if the retention period has passed.")
		redshift_batchModifyClusterSnapshotsCmd.Flags().String("manual-snapshot-retention-period", "", "The number of days that a manual snapshot is retained.")
		redshift_batchModifyClusterSnapshotsCmd.Flags().Bool("no-force", false, "A boolean value indicating whether to override an exception if the retention period has passed.")
		redshift_batchModifyClusterSnapshotsCmd.Flags().String("snapshot-identifier-list", "", "A list of snapshot identifiers you want to modify.")
		redshift_batchModifyClusterSnapshotsCmd.Flag("no-force").Hidden = true
		redshift_batchModifyClusterSnapshotsCmd.MarkFlagRequired("snapshot-identifier-list")
	})
	redshiftCmd.AddCommand(redshift_batchModifyClusterSnapshotsCmd)
}
