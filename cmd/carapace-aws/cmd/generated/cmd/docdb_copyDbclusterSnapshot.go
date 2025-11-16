package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_copyDbclusterSnapshotCmd = &cobra.Command{
	Use:   "copy-dbcluster-snapshot",
	Short: "Copies a snapshot of a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_copyDbclusterSnapshotCmd).Standalone()

	docdb_copyDbclusterSnapshotCmd.Flags().String("copy-tags", "", "Set to `true` to copy all tags from the source cluster snapshot to the target cluster snapshot, and otherwise `false`.")
	docdb_copyDbclusterSnapshotCmd.Flags().String("kms-key-id", "", "The KMS key ID for an encrypted cluster snapshot.")
	docdb_copyDbclusterSnapshotCmd.Flags().String("pre-signed-url", "", "The URL that contains a Signature Version 4 signed request for the`CopyDBClusterSnapshot` API action in the Amazon Web Services Region that contains the source cluster snapshot to copy.")
	docdb_copyDbclusterSnapshotCmd.Flags().String("source-dbcluster-snapshot-identifier", "", "The identifier of the cluster snapshot to copy.")
	docdb_copyDbclusterSnapshotCmd.Flags().String("tags", "", "The tags to be assigned to the cluster snapshot.")
	docdb_copyDbclusterSnapshotCmd.Flags().String("target-dbcluster-snapshot-identifier", "", "The identifier of the new cluster snapshot to create from the source cluster snapshot.")
	docdb_copyDbclusterSnapshotCmd.MarkFlagRequired("source-dbcluster-snapshot-identifier")
	docdb_copyDbclusterSnapshotCmd.MarkFlagRequired("target-dbcluster-snapshot-identifier")
	docdbCmd.AddCommand(docdb_copyDbclusterSnapshotCmd)
}
