package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifySnapshotCopyRetentionPeriodCmd = &cobra.Command{
	Use:   "modify-snapshot-copy-retention-period",
	Short: "Modifies the number of days to retain snapshots in the destination Amazon Web Services Region after they are copied from the source Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifySnapshotCopyRetentionPeriodCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_modifySnapshotCopyRetentionPeriodCmd).Standalone()

		redshift_modifySnapshotCopyRetentionPeriodCmd.Flags().String("cluster-identifier", "", "The unique identifier of the cluster for which you want to change the retention period for either automated or manual snapshots that are copied to a destination Amazon Web Services Region.")
		redshift_modifySnapshotCopyRetentionPeriodCmd.Flags().Bool("manual", false, "Indicates whether to apply the snapshot retention period to newly copied manual snapshots instead of automated snapshots.")
		redshift_modifySnapshotCopyRetentionPeriodCmd.Flags().Bool("no-manual", false, "Indicates whether to apply the snapshot retention period to newly copied manual snapshots instead of automated snapshots.")
		redshift_modifySnapshotCopyRetentionPeriodCmd.Flags().String("retention-period", "", "The number of days to retain automated snapshots in the destination Amazon Web Services Region after they are copied from the source Amazon Web Services Region.")
		redshift_modifySnapshotCopyRetentionPeriodCmd.MarkFlagRequired("cluster-identifier")
		redshift_modifySnapshotCopyRetentionPeriodCmd.Flag("no-manual").Hidden = true
		redshift_modifySnapshotCopyRetentionPeriodCmd.MarkFlagRequired("retention-period")
	})
	redshiftCmd.AddCommand(redshift_modifySnapshotCopyRetentionPeriodCmd)
}
