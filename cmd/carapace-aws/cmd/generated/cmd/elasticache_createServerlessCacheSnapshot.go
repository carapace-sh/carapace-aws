package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_createServerlessCacheSnapshotCmd = &cobra.Command{
	Use:   "create-serverless-cache-snapshot",
	Short: "This API creates a copy of an entire ServerlessCache at a specific moment in time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_createServerlessCacheSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_createServerlessCacheSnapshotCmd).Standalone()

		elasticache_createServerlessCacheSnapshotCmd.Flags().String("kms-key-id", "", "The ID of the KMS key used to encrypt the snapshot.")
		elasticache_createServerlessCacheSnapshotCmd.Flags().String("serverless-cache-name", "", "The name of an existing serverless cache.")
		elasticache_createServerlessCacheSnapshotCmd.Flags().String("serverless-cache-snapshot-name", "", "The name for the snapshot being created.")
		elasticache_createServerlessCacheSnapshotCmd.Flags().String("tags", "", "A list of tags to be added to the snapshot resource.")
		elasticache_createServerlessCacheSnapshotCmd.MarkFlagRequired("serverless-cache-name")
		elasticache_createServerlessCacheSnapshotCmd.MarkFlagRequired("serverless-cache-snapshot-name")
	})
	elasticacheCmd.AddCommand(elasticache_createServerlessCacheSnapshotCmd)
}
