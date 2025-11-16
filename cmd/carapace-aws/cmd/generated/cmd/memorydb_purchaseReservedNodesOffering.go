package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_purchaseReservedNodesOfferingCmd = &cobra.Command{
	Use:   "purchase-reserved-nodes-offering",
	Short: "Allows you to purchase a reserved node offering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_purchaseReservedNodesOfferingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_purchaseReservedNodesOfferingCmd).Standalone()

		memorydb_purchaseReservedNodesOfferingCmd.Flags().String("node-count", "", "The number of node instances to reserve.")
		memorydb_purchaseReservedNodesOfferingCmd.Flags().String("reservation-id", "", "A customer-specified identifier to track this reservation.")
		memorydb_purchaseReservedNodesOfferingCmd.Flags().String("reserved-nodes-offering-id", "", "The ID of the reserved node offering to purchase.")
		memorydb_purchaseReservedNodesOfferingCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
		memorydb_purchaseReservedNodesOfferingCmd.MarkFlagRequired("reserved-nodes-offering-id")
	})
	memorydbCmd.AddCommand(memorydb_purchaseReservedNodesOfferingCmd)
}
