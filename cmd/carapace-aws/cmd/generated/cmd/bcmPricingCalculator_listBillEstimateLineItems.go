package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_listBillEstimateLineItemsCmd = &cobra.Command{
	Use:   "list-bill-estimate-line-items",
	Short: "Lists the line items associated with a bill estimate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_listBillEstimateLineItemsCmd).Standalone()

	bcmPricingCalculator_listBillEstimateLineItemsCmd.Flags().String("bill-estimate-id", "", "The unique identifier of the bill estimate to list line items for.")
	bcmPricingCalculator_listBillEstimateLineItemsCmd.Flags().String("filters", "", "Filters to apply to the list of line items.")
	bcmPricingCalculator_listBillEstimateLineItemsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	bcmPricingCalculator_listBillEstimateLineItemsCmd.Flags().String("next-token", "", "A token to retrieve the next page of results.")
	bcmPricingCalculator_listBillEstimateLineItemsCmd.MarkFlagRequired("bill-estimate-id")
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_listBillEstimateLineItemsCmd)
}
