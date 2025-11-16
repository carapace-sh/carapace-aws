package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_listBillEstimateInputCommitmentModificationsCmd = &cobra.Command{
	Use:   "list-bill-estimate-input-commitment-modifications",
	Short: "Lists the input commitment modifications associated with a bill estimate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_listBillEstimateInputCommitmentModificationsCmd).Standalone()

	bcmPricingCalculator_listBillEstimateInputCommitmentModificationsCmd.Flags().String("bill-estimate-id", "", "The unique identifier of the bill estimate to list input commitment modifications for.")
	bcmPricingCalculator_listBillEstimateInputCommitmentModificationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	bcmPricingCalculator_listBillEstimateInputCommitmentModificationsCmd.Flags().String("next-token", "", "A token to retrieve the next page of results.")
	bcmPricingCalculator_listBillEstimateInputCommitmentModificationsCmd.MarkFlagRequired("bill-estimate-id")
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_listBillEstimateInputCommitmentModificationsCmd)
}
