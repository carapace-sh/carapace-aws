package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_purchaseReservedCacheNodesOfferingCmd = &cobra.Command{
	Use:   "purchase-reserved-cache-nodes-offering",
	Short: "Allows you to purchase a reserved cache node offering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_purchaseReservedCacheNodesOfferingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_purchaseReservedCacheNodesOfferingCmd).Standalone()

		elasticache_purchaseReservedCacheNodesOfferingCmd.Flags().String("cache-node-count", "", "The number of cache node instances to reserve.")
		elasticache_purchaseReservedCacheNodesOfferingCmd.Flags().String("reserved-cache-node-id", "", "A customer-specified identifier to track this reservation.")
		elasticache_purchaseReservedCacheNodesOfferingCmd.Flags().String("reserved-cache-nodes-offering-id", "", "The ID of the reserved cache node offering to purchase.")
		elasticache_purchaseReservedCacheNodesOfferingCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
		elasticache_purchaseReservedCacheNodesOfferingCmd.MarkFlagRequired("reserved-cache-nodes-offering-id")
	})
	elasticacheCmd.AddCommand(elasticache_purchaseReservedCacheNodesOfferingCmd)
}
