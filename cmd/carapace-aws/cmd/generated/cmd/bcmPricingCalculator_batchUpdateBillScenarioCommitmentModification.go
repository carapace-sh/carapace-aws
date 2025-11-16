package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_batchUpdateBillScenarioCommitmentModificationCmd = &cobra.Command{
	Use:   "batch-update-bill-scenario-commitment-modification",
	Short: "Update a newly added or existing commitment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_batchUpdateBillScenarioCommitmentModificationCmd).Standalone()

	bcmPricingCalculator_batchUpdateBillScenarioCommitmentModificationCmd.Flags().String("bill-scenario-id", "", "The ID of the Bill Scenario for which you want to modify the commitment group of a modeled commitment.")
	bcmPricingCalculator_batchUpdateBillScenarioCommitmentModificationCmd.Flags().String("commitment-modifications", "", "List of commitments that you want to update in a Bill Scenario.")
	bcmPricingCalculator_batchUpdateBillScenarioCommitmentModificationCmd.MarkFlagRequired("bill-scenario-id")
	bcmPricingCalculator_batchUpdateBillScenarioCommitmentModificationCmd.MarkFlagRequired("commitment-modifications")
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_batchUpdateBillScenarioCommitmentModificationCmd)
}
