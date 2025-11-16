package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_purchaseReservedDbinstancesOfferingCmd = &cobra.Command{
	Use:   "purchase-reserved-dbinstances-offering",
	Short: "Purchases a reserved DB instance offering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_purchaseReservedDbinstancesOfferingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_purchaseReservedDbinstancesOfferingCmd).Standalone()

		rds_purchaseReservedDbinstancesOfferingCmd.Flags().String("dbinstance-count", "", "The number of instances to reserve.")
		rds_purchaseReservedDbinstancesOfferingCmd.Flags().String("reserved-dbinstance-id", "", "Customer-specified identifier to track this reservation.")
		rds_purchaseReservedDbinstancesOfferingCmd.Flags().String("reserved-dbinstances-offering-id", "", "The ID of the Reserved DB instance offering to purchase.")
		rds_purchaseReservedDbinstancesOfferingCmd.Flags().String("tags", "", "")
		rds_purchaseReservedDbinstancesOfferingCmd.MarkFlagRequired("reserved-dbinstances-offering-id")
	})
	rdsCmd.AddCommand(rds_purchaseReservedDbinstancesOfferingCmd)
}
