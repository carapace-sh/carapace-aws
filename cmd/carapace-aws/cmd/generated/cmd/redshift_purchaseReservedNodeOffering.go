package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_purchaseReservedNodeOfferingCmd = &cobra.Command{
	Use:   "purchase-reserved-node-offering",
	Short: "Allows you to purchase reserved nodes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_purchaseReservedNodeOfferingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_purchaseReservedNodeOfferingCmd).Standalone()

		redshift_purchaseReservedNodeOfferingCmd.Flags().String("node-count", "", "The number of reserved nodes that you want to purchase.")
		redshift_purchaseReservedNodeOfferingCmd.Flags().String("reserved-node-offering-id", "", "The unique identifier of the reserved node offering you want to purchase.")
		redshift_purchaseReservedNodeOfferingCmd.MarkFlagRequired("reserved-node-offering-id")
	})
	redshiftCmd.AddCommand(redshift_purchaseReservedNodeOfferingCmd)
}
