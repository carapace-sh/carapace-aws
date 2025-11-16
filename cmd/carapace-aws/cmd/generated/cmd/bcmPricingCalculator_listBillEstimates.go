package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_listBillEstimatesCmd = &cobra.Command{
	Use:   "list-bill-estimates",
	Short: "Lists all bill estimates for the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_listBillEstimatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_listBillEstimatesCmd).Standalone()

		bcmPricingCalculator_listBillEstimatesCmd.Flags().String("created-at-filter", "", "Filter bill estimates based on the creation date.")
		bcmPricingCalculator_listBillEstimatesCmd.Flags().String("expires-at-filter", "", "Filter bill estimates based on the expiration date.")
		bcmPricingCalculator_listBillEstimatesCmd.Flags().String("filters", "", "Filters to apply to the list of bill estimates.")
		bcmPricingCalculator_listBillEstimatesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		bcmPricingCalculator_listBillEstimatesCmd.Flags().String("next-token", "", "A token to retrieve the next page of results.")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_listBillEstimatesCmd)
}
