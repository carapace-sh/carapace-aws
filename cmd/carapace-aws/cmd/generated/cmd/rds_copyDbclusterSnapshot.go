package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_copyDbclusterSnapshotCmd = &cobra.Command{
	Use:   "copy-dbcluster-snapshot",
	Short: "Copies a snapshot of a DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_copyDbclusterSnapshotCmd).Standalone()

	rds_copyDbclusterSnapshotCmd.Flags().String("copy-tags", "", "Specifies whether to copy all tags from the source DB cluster snapshot to the target DB cluster snapshot.")
	rds_copyDbclusterSnapshotCmd.Flags().String("kms-key-id", "", "The Amazon Web Services KMS key identifier for an encrypted DB cluster snapshot.")
	rds_copyDbclusterSnapshotCmd.Flags().String("pre-signed-url", "", "When you are copying a DB cluster snapshot from one Amazon Web Services GovCloud (US) Region to another, the URL that contains a Signature Version 4 signed request for the `CopyDBClusterSnapshot` API operation in the Amazon Web Services Region that contains the source DB cluster snapshot to copy.")
	rds_copyDbclusterSnapshotCmd.Flags().String("source-dbcluster-snapshot-identifier", "", "The identifier of the DB cluster snapshot to copy.")
	rds_copyDbclusterSnapshotCmd.Flags().String("tags", "", "")
	rds_copyDbclusterSnapshotCmd.Flags().String("target-dbcluster-snapshot-identifier", "", "The identifier of the new DB cluster snapshot to create from the source DB cluster snapshot.")
	rds_copyDbclusterSnapshotCmd.MarkFlagRequired("source-dbcluster-snapshot-identifier")
	rds_copyDbclusterSnapshotCmd.MarkFlagRequired("target-dbcluster-snapshot-identifier")
	rdsCmd.AddCommand(rds_copyDbclusterSnapshotCmd)
}
