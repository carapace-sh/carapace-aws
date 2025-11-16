package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_purchaseReservedInstanceOfferingCmd = &cobra.Command{
	Use:   "purchase-reserved-instance-offering",
	Short: "Allows you to purchase Amazon OpenSearch Service Reserved Instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_purchaseReservedInstanceOfferingCmd).Standalone()

	opensearch_purchaseReservedInstanceOfferingCmd.Flags().String("instance-count", "", "The number of OpenSearch instances to reserve.")
	opensearch_purchaseReservedInstanceOfferingCmd.Flags().String("reservation-name", "", "A customer-specified identifier to track this reservation.")
	opensearch_purchaseReservedInstanceOfferingCmd.Flags().String("reserved-instance-offering-id", "", "The ID of the Reserved Instance offering to purchase.")
	opensearch_purchaseReservedInstanceOfferingCmd.MarkFlagRequired("reservation-name")
	opensearch_purchaseReservedInstanceOfferingCmd.MarkFlagRequired("reserved-instance-offering-id")
	opensearchCmd.AddCommand(opensearch_purchaseReservedInstanceOfferingCmd)
}
