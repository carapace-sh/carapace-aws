package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_listBillScenarioCommitmentModificationsCmd = &cobra.Command{
	Use:   "list-bill-scenario-commitment-modifications",
	Short: "Lists the commitment modifications associated with a bill scenario.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_listBillScenarioCommitmentModificationsCmd).Standalone()

	bcmPricingCalculator_listBillScenarioCommitmentModificationsCmd.Flags().String("bill-scenario-id", "", "The unique identifier of the bill scenario to list commitment modifications for.")
	bcmPricingCalculator_listBillScenarioCommitmentModificationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	bcmPricingCalculator_listBillScenarioCommitmentModificationsCmd.Flags().String("next-token", "", "A token to retrieve the next page of results.")
	bcmPricingCalculator_listBillScenarioCommitmentModificationsCmd.MarkFlagRequired("bill-scenario-id")
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_listBillScenarioCommitmentModificationsCmd)
}
