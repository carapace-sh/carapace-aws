package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeReservedCacheNodesOfferingsCmd = &cobra.Command{
	Use:   "describe-reserved-cache-nodes-offerings",
	Short: "Lists available reserved cache node offerings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeReservedCacheNodesOfferingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_describeReservedCacheNodesOfferingsCmd).Standalone()

		elasticache_describeReservedCacheNodesOfferingsCmd.Flags().String("cache-node-type", "", "The cache node type filter value.")
		elasticache_describeReservedCacheNodesOfferingsCmd.Flags().String("duration", "", "Duration filter value, specified in years or seconds.")
		elasticache_describeReservedCacheNodesOfferingsCmd.Flags().String("marker", "", "An optional marker returned from a prior request.")
		elasticache_describeReservedCacheNodesOfferingsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		elasticache_describeReservedCacheNodesOfferingsCmd.Flags().String("offering-type", "", "The offering type filter value.")
		elasticache_describeReservedCacheNodesOfferingsCmd.Flags().String("product-description", "", "The product description filter value.")
		elasticache_describeReservedCacheNodesOfferingsCmd.Flags().String("reserved-cache-nodes-offering-id", "", "The offering identifier filter value.")
	})
	elasticacheCmd.AddCommand(elasticache_describeReservedCacheNodesOfferingsCmd)
}
