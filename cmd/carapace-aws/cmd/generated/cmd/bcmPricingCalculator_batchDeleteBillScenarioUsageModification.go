package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_batchDeleteBillScenarioUsageModificationCmd = &cobra.Command{
	Use:   "batch-delete-bill-scenario-usage-modification",
	Short: "Delete usage that you have created in a Bill Scenario.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_batchDeleteBillScenarioUsageModificationCmd).Standalone()

	bcmPricingCalculator_batchDeleteBillScenarioUsageModificationCmd.Flags().String("bill-scenario-id", "", "The ID of the Bill Scenario for which you want to delete the modeled usage.")
	bcmPricingCalculator_batchDeleteBillScenarioUsageModificationCmd.Flags().String("ids", "", "List of usage that you want to delete from the Bill Scenario.")
	bcmPricingCalculator_batchDeleteBillScenarioUsageModificationCmd.MarkFlagRequired("bill-scenario-id")
	bcmPricingCalculator_batchDeleteBillScenarioUsageModificationCmd.MarkFlagRequired("ids")
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_batchDeleteBillScenarioUsageModificationCmd)
}
