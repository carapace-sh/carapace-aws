package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_copySnapshotCmd = &cobra.Command{
	Use:   "copy-snapshot",
	Short: "Makes a copy of an existing snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_copySnapshotCmd).Standalone()

	memorydb_copySnapshotCmd.Flags().String("kms-key-id", "", "The ID of the KMS key used to encrypt the target snapshot.")
	memorydb_copySnapshotCmd.Flags().String("source-snapshot-name", "", "The name of an existing snapshot from which to make a copy.")
	memorydb_copySnapshotCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
	memorydb_copySnapshotCmd.Flags().String("target-bucket", "", "The Amazon S3 bucket to which the snapshot is exported.")
	memorydb_copySnapshotCmd.Flags().String("target-snapshot-name", "", "A name for the snapshot copy.")
	memorydb_copySnapshotCmd.MarkFlagRequired("source-snapshot-name")
	memorydb_copySnapshotCmd.MarkFlagRequired("target-snapshot-name")
	memorydbCmd.AddCommand(memorydb_copySnapshotCmd)
}
