package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_getBillScenarioCmd = &cobra.Command{
	Use:   "get-bill-scenario",
	Short: "Retrieves details of a specific bill scenario.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_getBillScenarioCmd).Standalone()

	bcmPricingCalculator_getBillScenarioCmd.Flags().String("identifier", "", "The unique identifier of the bill scenario to retrieve.")
	bcmPricingCalculator_getBillScenarioCmd.MarkFlagRequired("identifier")
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_getBillScenarioCmd)
}
