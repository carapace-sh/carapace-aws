package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_purchaseOfferingCmd = &cobra.Command{
	Use:   "purchase-offering",
	Short: "Immediately purchases offerings for an AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_purchaseOfferingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_purchaseOfferingCmd).Standalone()

		devicefarm_purchaseOfferingCmd.Flags().String("offering-id", "", "The ID of the offering.")
		devicefarm_purchaseOfferingCmd.Flags().String("offering-promotion-id", "", "The ID of the offering promotion to be applied to the purchase.")
		devicefarm_purchaseOfferingCmd.Flags().String("quantity", "", "The number of device slots to purchase in an offering request.")
		devicefarm_purchaseOfferingCmd.MarkFlagRequired("offering-id")
		devicefarm_purchaseOfferingCmd.MarkFlagRequired("quantity")
	})
	devicefarmCmd.AddCommand(devicefarm_purchaseOfferingCmd)
}
