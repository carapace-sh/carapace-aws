package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_batchCreateBillScenarioCommitmentModificationCmd = &cobra.Command{
	Use:   "batch-create-bill-scenario-commitment-modification",
	Short: "Create Compute Savings Plans, EC2 Instance Savings Plans, or EC2 Reserved Instances commitments that you want to model in a Bill Scenario.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_batchCreateBillScenarioCommitmentModificationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_batchCreateBillScenarioCommitmentModificationCmd).Standalone()

		bcmPricingCalculator_batchCreateBillScenarioCommitmentModificationCmd.Flags().String("bill-scenario-id", "", "The ID of the Bill Scenario for which you want to create the modeled commitment.")
		bcmPricingCalculator_batchCreateBillScenarioCommitmentModificationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		bcmPricingCalculator_batchCreateBillScenarioCommitmentModificationCmd.Flags().String("commitment-modifications", "", "List of commitments that you want to model in the Bill Scenario.")
		bcmPricingCalculator_batchCreateBillScenarioCommitmentModificationCmd.MarkFlagRequired("bill-scenario-id")
		bcmPricingCalculator_batchCreateBillScenarioCommitmentModificationCmd.MarkFlagRequired("commitment-modifications")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_batchCreateBillScenarioCommitmentModificationCmd)
}
