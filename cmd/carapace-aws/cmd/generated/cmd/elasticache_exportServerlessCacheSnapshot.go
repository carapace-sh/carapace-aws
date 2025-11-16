package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_exportServerlessCacheSnapshotCmd = &cobra.Command{
	Use:   "export-serverless-cache-snapshot",
	Short: "Provides the functionality to export the serverless cache snapshot data to Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_exportServerlessCacheSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_exportServerlessCacheSnapshotCmd).Standalone()

		elasticache_exportServerlessCacheSnapshotCmd.Flags().String("s3-bucket-name", "", "Name of the Amazon S3 bucket to export the snapshot to.")
		elasticache_exportServerlessCacheSnapshotCmd.Flags().String("serverless-cache-snapshot-name", "", "The identifier of the serverless cache snapshot to be exported to S3.")
		elasticache_exportServerlessCacheSnapshotCmd.MarkFlagRequired("s3-bucket-name")
		elasticache_exportServerlessCacheSnapshotCmd.MarkFlagRequired("serverless-cache-snapshot-name")
	})
	elasticacheCmd.AddCommand(elasticache_exportServerlessCacheSnapshotCmd)
}
