package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_copyServerlessCacheSnapshotCmd = &cobra.Command{
	Use:   "copy-serverless-cache-snapshot",
	Short: "Creates a copy of an existing serverless cache’s snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_copyServerlessCacheSnapshotCmd).Standalone()

	elasticache_copyServerlessCacheSnapshotCmd.Flags().String("kms-key-id", "", "The identifier of the KMS key used to encrypt the target snapshot.")
	elasticache_copyServerlessCacheSnapshotCmd.Flags().String("source-serverless-cache-snapshot-name", "", "The identifier of the existing serverless cache’s snapshot to be copied.")
	elasticache_copyServerlessCacheSnapshotCmd.Flags().String("tags", "", "A list of tags to be added to the target snapshot resource.")
	elasticache_copyServerlessCacheSnapshotCmd.Flags().String("target-serverless-cache-snapshot-name", "", "The identifier for the snapshot to be created.")
	elasticache_copyServerlessCacheSnapshotCmd.MarkFlagRequired("source-serverless-cache-snapshot-name")
	elasticache_copyServerlessCacheSnapshotCmd.MarkFlagRequired("target-serverless-cache-snapshot-name")
	elasticacheCmd.AddCommand(elasticache_copyServerlessCacheSnapshotCmd)
}
