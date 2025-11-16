package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeServerlessCacheSnapshotsCmd = &cobra.Command{
	Use:   "describe-serverless-cache-snapshots",
	Short: "Returns information about serverless cache snapshots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeServerlessCacheSnapshotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_describeServerlessCacheSnapshotsCmd).Standalone()

		elasticache_describeServerlessCacheSnapshotsCmd.Flags().String("max-results", "", "The maximum number of records to include in the response.")
		elasticache_describeServerlessCacheSnapshotsCmd.Flags().String("next-token", "", "An optional marker returned from a prior request to support pagination of results from this operation.")
		elasticache_describeServerlessCacheSnapshotsCmd.Flags().String("serverless-cache-name", "", "The identifier of serverless cache.")
		elasticache_describeServerlessCacheSnapshotsCmd.Flags().String("serverless-cache-snapshot-name", "", "The identifier of the serverless cache’s snapshot.")
		elasticache_describeServerlessCacheSnapshotsCmd.Flags().String("snapshot-type", "", "The type of snapshot that is being described.")
	})
	elasticacheCmd.AddCommand(elasticache_describeServerlessCacheSnapshotsCmd)
}
