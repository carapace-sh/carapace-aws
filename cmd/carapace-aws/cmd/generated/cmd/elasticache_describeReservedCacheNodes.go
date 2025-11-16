package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeReservedCacheNodesCmd = &cobra.Command{
	Use:   "describe-reserved-cache-nodes",
	Short: "Returns information about reserved cache nodes for this account, or about a specified reserved cache node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeReservedCacheNodesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_describeReservedCacheNodesCmd).Standalone()

		elasticache_describeReservedCacheNodesCmd.Flags().String("cache-node-type", "", "The cache node type filter value.")
		elasticache_describeReservedCacheNodesCmd.Flags().String("duration", "", "The duration filter value, specified in years or seconds.")
		elasticache_describeReservedCacheNodesCmd.Flags().String("marker", "", "An optional marker returned from a prior request.")
		elasticache_describeReservedCacheNodesCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		elasticache_describeReservedCacheNodesCmd.Flags().String("offering-type", "", "The offering type filter value.")
		elasticache_describeReservedCacheNodesCmd.Flags().String("product-description", "", "The product description filter value.")
		elasticache_describeReservedCacheNodesCmd.Flags().String("reserved-cache-node-id", "", "The reserved cache node identifier filter value.")
		elasticache_describeReservedCacheNodesCmd.Flags().String("reserved-cache-nodes-offering-id", "", "The offering identifier filter value.")
	})
	elasticacheCmd.AddCommand(elasticache_describeReservedCacheNodesCmd)
}
