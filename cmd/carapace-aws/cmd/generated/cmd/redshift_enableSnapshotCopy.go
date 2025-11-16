package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_enableSnapshotCopyCmd = &cobra.Command{
	Use:   "enable-snapshot-copy",
	Short: "Enables the automatic copy of snapshots from one region to another region for a specified cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_enableSnapshotCopyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_enableSnapshotCopyCmd).Standalone()

		redshift_enableSnapshotCopyCmd.Flags().String("cluster-identifier", "", "The unique identifier of the source cluster to copy snapshots from.")
		redshift_enableSnapshotCopyCmd.Flags().String("destination-region", "", "The destination Amazon Web Services Region that you want to copy snapshots to.")
		redshift_enableSnapshotCopyCmd.Flags().String("manual-snapshot-retention-period", "", "The number of days to retain newly copied snapshots in the destination Amazon Web Services Region after they are copied from the source Amazon Web Services Region.")
		redshift_enableSnapshotCopyCmd.Flags().String("retention-period", "", "The number of days to retain automated snapshots in the destination region after they are copied from the source region.")
		redshift_enableSnapshotCopyCmd.Flags().String("snapshot-copy-grant-name", "", "The name of the snapshot copy grant to use when snapshots of an Amazon Web Services KMS-encrypted cluster are copied to the destination region.")
		redshift_enableSnapshotCopyCmd.MarkFlagRequired("cluster-identifier")
		redshift_enableSnapshotCopyCmd.MarkFlagRequired("destination-region")
	})
	redshiftCmd.AddCommand(redshift_enableSnapshotCopyCmd)
}
