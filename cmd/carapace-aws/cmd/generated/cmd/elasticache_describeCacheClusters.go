package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeCacheClustersCmd = &cobra.Command{
	Use:   "describe-cache-clusters",
	Short: "Returns information about all provisioned clusters if no cluster identifier is specified, or about a specific cache cluster if a cluster identifier is supplied.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeCacheClustersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_describeCacheClustersCmd).Standalone()

		elasticache_describeCacheClustersCmd.Flags().String("cache-cluster-id", "", "The user-supplied cluster identifier.")
		elasticache_describeCacheClustersCmd.Flags().String("marker", "", "An optional marker returned from a prior request.")
		elasticache_describeCacheClustersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		elasticache_describeCacheClustersCmd.Flags().String("show-cache-clusters-not-in-replication-groups", "", "An optional flag that can be included in the `DescribeCacheCluster` request to show only nodes (API/CLI: clusters) that are not members of a replication group.")
		elasticache_describeCacheClustersCmd.Flags().String("show-cache-node-info", "", "An optional flag that can be included in the `DescribeCacheCluster` request to retrieve information about the individual cache nodes.")
	})
	elasticacheCmd.AddCommand(elasticache_describeCacheClustersCmd)
}
