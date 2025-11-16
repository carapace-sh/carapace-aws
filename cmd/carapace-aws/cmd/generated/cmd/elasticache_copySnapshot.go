package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_copySnapshotCmd = &cobra.Command{
	Use:   "copy-snapshot",
	Short: "Makes a copy of an existing snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_copySnapshotCmd).Standalone()

	elasticache_copySnapshotCmd.Flags().String("kms-key-id", "", "The ID of the KMS key used to encrypt the target snapshot.")
	elasticache_copySnapshotCmd.Flags().String("source-snapshot-name", "", "The name of an existing snapshot from which to make a copy.")
	elasticache_copySnapshotCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
	elasticache_copySnapshotCmd.Flags().String("target-bucket", "", "The Amazon S3 bucket to which the snapshot is exported.")
	elasticache_copySnapshotCmd.Flags().String("target-snapshot-name", "", "A name for the snapshot copy.")
	elasticache_copySnapshotCmd.MarkFlagRequired("source-snapshot-name")
	elasticache_copySnapshotCmd.MarkFlagRequired("target-snapshot-name")
	elasticacheCmd.AddCommand(elasticache_copySnapshotCmd)
}
