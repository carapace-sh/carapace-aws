package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_createBillScenarioCmd = &cobra.Command{
	Use:   "create-bill-scenario",
	Short: "Creates a new bill scenario to model potential changes to Amazon Web Services usage and costs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_createBillScenarioCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_createBillScenarioCmd).Standalone()

		bcmPricingCalculator_createBillScenarioCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure idempotency of the request.")
		bcmPricingCalculator_createBillScenarioCmd.Flags().String("name", "", "A descriptive name for the bill scenario.")
		bcmPricingCalculator_createBillScenarioCmd.Flags().String("tags", "", "The tags to apply to the bill scenario.")
		bcmPricingCalculator_createBillScenarioCmd.MarkFlagRequired("name")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_createBillScenarioCmd)
}
