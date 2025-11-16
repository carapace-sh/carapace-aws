package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_deleteBillEstimateCmd = &cobra.Command{
	Use:   "delete-bill-estimate",
	Short: "Deletes an existing bill estimate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_deleteBillEstimateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_deleteBillEstimateCmd).Standalone()

		bcmPricingCalculator_deleteBillEstimateCmd.Flags().String("identifier", "", "The unique identifier of the bill estimate to delete.")
		bcmPricingCalculator_deleteBillEstimateCmd.MarkFlagRequired("identifier")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_deleteBillEstimateCmd)
}
