package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_updateBillEstimateCmd = &cobra.Command{
	Use:   "update-bill-estimate",
	Short: "Updates an existing bill estimate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_updateBillEstimateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_updateBillEstimateCmd).Standalone()

		bcmPricingCalculator_updateBillEstimateCmd.Flags().String("expires-at", "", "The new expiration date for the bill estimate.")
		bcmPricingCalculator_updateBillEstimateCmd.Flags().String("identifier", "", "The unique identifier of the bill estimate to update.")
		bcmPricingCalculator_updateBillEstimateCmd.Flags().String("name", "", "The new name for the bill estimate.")
		bcmPricingCalculator_updateBillEstimateCmd.MarkFlagRequired("identifier")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_updateBillEstimateCmd)
}
