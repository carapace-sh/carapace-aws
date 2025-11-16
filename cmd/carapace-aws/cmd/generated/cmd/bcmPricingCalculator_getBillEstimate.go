package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_getBillEstimateCmd = &cobra.Command{
	Use:   "get-bill-estimate",
	Short: "Retrieves details of a specific bill estimate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_getBillEstimateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_getBillEstimateCmd).Standalone()

		bcmPricingCalculator_getBillEstimateCmd.Flags().String("identifier", "", "The unique identifier of the bill estimate to retrieve.")
		bcmPricingCalculator_getBillEstimateCmd.MarkFlagRequired("identifier")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_getBillEstimateCmd)
}
