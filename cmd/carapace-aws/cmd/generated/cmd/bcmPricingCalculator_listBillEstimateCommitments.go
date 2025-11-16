package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_listBillEstimateCommitmentsCmd = &cobra.Command{
	Use:   "list-bill-estimate-commitments",
	Short: "Lists the commitments associated with a bill estimate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_listBillEstimateCommitmentsCmd).Standalone()

	bcmPricingCalculator_listBillEstimateCommitmentsCmd.Flags().String("bill-estimate-id", "", "The unique identifier of the bill estimate to list commitments for.")
	bcmPricingCalculator_listBillEstimateCommitmentsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	bcmPricingCalculator_listBillEstimateCommitmentsCmd.Flags().String("next-token", "", "A token to retrieve the next page of results.")
	bcmPricingCalculator_listBillEstimateCommitmentsCmd.MarkFlagRequired("bill-estimate-id")
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_listBillEstimateCommitmentsCmd)
}
