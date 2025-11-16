package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_batchCreateBillScenarioUsageModificationCmd = &cobra.Command{
	Use:   "batch-create-bill-scenario-usage-modification",
	Short: "Create Amazon Web Services service usage that you want to model in a Bill Scenario.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_batchCreateBillScenarioUsageModificationCmd).Standalone()

	bcmPricingCalculator_batchCreateBillScenarioUsageModificationCmd.Flags().String("bill-scenario-id", "", "The ID of the Bill Scenario for which you want to create the modeled usage.")
	bcmPricingCalculator_batchCreateBillScenarioUsageModificationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	bcmPricingCalculator_batchCreateBillScenarioUsageModificationCmd.Flags().String("usage-modifications", "", "List of usage that you want to model in the Bill Scenario.")
	bcmPricingCalculator_batchCreateBillScenarioUsageModificationCmd.MarkFlagRequired("bill-scenario-id")
	bcmPricingCalculator_batchCreateBillScenarioUsageModificationCmd.MarkFlagRequired("usage-modifications")
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_batchCreateBillScenarioUsageModificationCmd)
}
