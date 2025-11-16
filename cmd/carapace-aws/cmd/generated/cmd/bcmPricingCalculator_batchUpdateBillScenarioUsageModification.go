package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_batchUpdateBillScenarioUsageModificationCmd = &cobra.Command{
	Use:   "batch-update-bill-scenario-usage-modification",
	Short: "Update a newly added or existing usage lines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_batchUpdateBillScenarioUsageModificationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_batchUpdateBillScenarioUsageModificationCmd).Standalone()

		bcmPricingCalculator_batchUpdateBillScenarioUsageModificationCmd.Flags().String("bill-scenario-id", "", "The ID of the Bill Scenario for which you want to modify the usage lines.")
		bcmPricingCalculator_batchUpdateBillScenarioUsageModificationCmd.Flags().String("usage-modifications", "", "List of usage lines that you want to update in a Bill Scenario identified by the usage ID.")
		bcmPricingCalculator_batchUpdateBillScenarioUsageModificationCmd.MarkFlagRequired("bill-scenario-id")
		bcmPricingCalculator_batchUpdateBillScenarioUsageModificationCmd.MarkFlagRequired("usage-modifications")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_batchUpdateBillScenarioUsageModificationCmd)
}
