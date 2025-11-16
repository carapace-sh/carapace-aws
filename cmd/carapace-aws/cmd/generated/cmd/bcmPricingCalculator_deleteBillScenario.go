package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_deleteBillScenarioCmd = &cobra.Command{
	Use:   "delete-bill-scenario",
	Short: "Deletes an existing bill scenario.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_deleteBillScenarioCmd).Standalone()

	bcmPricingCalculator_deleteBillScenarioCmd.Flags().String("identifier", "", "The unique identifier of the bill scenario to delete.")
	bcmPricingCalculator_deleteBillScenarioCmd.MarkFlagRequired("identifier")
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_deleteBillScenarioCmd)
}
