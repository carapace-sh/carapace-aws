package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_copyDbclusterSnapshotCmd = &cobra.Command{
	Use:   "copy-dbcluster-snapshot",
	Short: "Copies a snapshot of a DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_copyDbclusterSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_copyDbclusterSnapshotCmd).Standalone()

		neptune_copyDbclusterSnapshotCmd.Flags().String("copy-tags", "", "True to copy all tags from the source DB cluster snapshot to the target DB cluster snapshot, and otherwise false.")
		neptune_copyDbclusterSnapshotCmd.Flags().String("kms-key-id", "", "The Amazon Amazon KMS key ID for an encrypted DB cluster snapshot.")
		neptune_copyDbclusterSnapshotCmd.Flags().String("pre-signed-url", "", "Not currently supported.")
		neptune_copyDbclusterSnapshotCmd.Flags().String("source-dbcluster-snapshot-identifier", "", "The identifier of the DB cluster snapshot to copy.")
		neptune_copyDbclusterSnapshotCmd.Flags().String("tags", "", "The tags to assign to the new DB cluster snapshot copy.")
		neptune_copyDbclusterSnapshotCmd.Flags().String("target-dbcluster-snapshot-identifier", "", "The identifier of the new DB cluster snapshot to create from the source DB cluster snapshot.")
		neptune_copyDbclusterSnapshotCmd.MarkFlagRequired("source-dbcluster-snapshot-identifier")
		neptune_copyDbclusterSnapshotCmd.MarkFlagRequired("target-dbcluster-snapshot-identifier")
	})
	neptuneCmd.AddCommand(neptune_copyDbclusterSnapshotCmd)
}
