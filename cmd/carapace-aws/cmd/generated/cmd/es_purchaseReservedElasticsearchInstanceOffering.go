package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_purchaseReservedElasticsearchInstanceOfferingCmd = &cobra.Command{
	Use:   "purchase-reserved-elasticsearch-instance-offering",
	Short: "Allows you to purchase reserved Elasticsearch instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_purchaseReservedElasticsearchInstanceOfferingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_purchaseReservedElasticsearchInstanceOfferingCmd).Standalone()

		es_purchaseReservedElasticsearchInstanceOfferingCmd.Flags().String("instance-count", "", "The number of Elasticsearch instances to reserve.")
		es_purchaseReservedElasticsearchInstanceOfferingCmd.Flags().String("reservation-name", "", "A customer-specified identifier to track this reservation.")
		es_purchaseReservedElasticsearchInstanceOfferingCmd.Flags().String("reserved-elasticsearch-instance-offering-id", "", "The ID of the reserved Elasticsearch instance offering to purchase.")
		es_purchaseReservedElasticsearchInstanceOfferingCmd.MarkFlagRequired("reservation-name")
		es_purchaseReservedElasticsearchInstanceOfferingCmd.MarkFlagRequired("reserved-elasticsearch-instance-offering-id")
	})
	esCmd.AddCommand(es_purchaseReservedElasticsearchInstanceOfferingCmd)
}
