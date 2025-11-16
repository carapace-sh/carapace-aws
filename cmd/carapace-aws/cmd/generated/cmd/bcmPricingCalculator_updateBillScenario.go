package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_updateBillScenarioCmd = &cobra.Command{
	Use:   "update-bill-scenario",
	Short: "Updates an existing bill scenario.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_updateBillScenarioCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_updateBillScenarioCmd).Standalone()

		bcmPricingCalculator_updateBillScenarioCmd.Flags().String("expires-at", "", "The new expiration date for the bill scenario.")
		bcmPricingCalculator_updateBillScenarioCmd.Flags().String("identifier", "", "The unique identifier of the bill scenario to update.")
		bcmPricingCalculator_updateBillScenarioCmd.Flags().String("name", "", "The new name for the bill scenario.")
		bcmPricingCalculator_updateBillScenarioCmd.MarkFlagRequired("identifier")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_updateBillScenarioCmd)
}
