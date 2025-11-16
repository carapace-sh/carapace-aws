package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_copyDbsnapshotCmd = &cobra.Command{
	Use:   "copy-dbsnapshot",
	Short: "Copies the specified DB snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_copyDbsnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_copyDbsnapshotCmd).Standalone()

		rds_copyDbsnapshotCmd.Flags().String("copy-option-group", "", "Specifies whether to copy the DB option group associated with the source DB snapshot to the target Amazon Web Services account and associate with the target DB snapshot.")
		rds_copyDbsnapshotCmd.Flags().String("copy-tags", "", "Specifies whether to copy all tags from the source DB snapshot to the target DB snapshot.")
		rds_copyDbsnapshotCmd.Flags().String("kms-key-id", "", "The Amazon Web Services KMS key identifier for an encrypted DB snapshot.")
		rds_copyDbsnapshotCmd.Flags().String("option-group-name", "", "The name of an option group to associate with the copy of the snapshot.")
		rds_copyDbsnapshotCmd.Flags().String("pre-signed-url", "", "When you are copying a snapshot from one Amazon Web Services GovCloud (US) Region to another, the URL that contains a Signature Version 4 signed request for the `CopyDBSnapshot` API operation in the source Amazon Web Services Region that contains the source DB snapshot to copy.")
		rds_copyDbsnapshotCmd.Flags().String("snapshot-availability-zone", "", "Specifies the name of the Availability Zone where RDS stores the DB snapshot.")
		rds_copyDbsnapshotCmd.Flags().String("snapshot-target", "", "Configures the location where RDS will store copied snapshots.")
		rds_copyDbsnapshotCmd.Flags().String("source-dbsnapshot-identifier", "", "The identifier for the source DB snapshot.")
		rds_copyDbsnapshotCmd.Flags().String("tags", "", "")
		rds_copyDbsnapshotCmd.Flags().String("target-custom-availability-zone", "", "The external custom Availability Zone (CAZ) identifier for the target CAZ.")
		rds_copyDbsnapshotCmd.Flags().String("target-dbsnapshot-identifier", "", "The identifier for the copy of the snapshot.")
		rds_copyDbsnapshotCmd.MarkFlagRequired("source-dbsnapshot-identifier")
		rds_copyDbsnapshotCmd.MarkFlagRequired("target-dbsnapshot-identifier")
	})
	rdsCmd.AddCommand(rds_copyDbsnapshotCmd)
}
