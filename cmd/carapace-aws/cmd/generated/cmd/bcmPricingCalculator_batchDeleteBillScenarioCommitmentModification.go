package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_batchDeleteBillScenarioCommitmentModificationCmd = &cobra.Command{
	Use:   "batch-delete-bill-scenario-commitment-modification",
	Short: "Delete commitment that you have created in a Bill Scenario.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_batchDeleteBillScenarioCommitmentModificationCmd).Standalone()

	bcmPricingCalculator_batchDeleteBillScenarioCommitmentModificationCmd.Flags().String("bill-scenario-id", "", "The ID of the Bill Scenario for which you want to delete the modeled commitment.")
	bcmPricingCalculator_batchDeleteBillScenarioCommitmentModificationCmd.Flags().String("ids", "", "List of commitments that you want to delete from the Bill Scenario.")
	bcmPricingCalculator_batchDeleteBillScenarioCommitmentModificationCmd.MarkFlagRequired("bill-scenario-id")
	bcmPricingCalculator_batchDeleteBillScenarioCommitmentModificationCmd.MarkFlagRequired("ids")
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_batchDeleteBillScenarioCommitmentModificationCmd)
}
