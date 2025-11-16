package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_listBillEstimateInputUsageModificationsCmd = &cobra.Command{
	Use:   "list-bill-estimate-input-usage-modifications",
	Short: "Lists the input usage modifications associated with a bill estimate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_listBillEstimateInputUsageModificationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_listBillEstimateInputUsageModificationsCmd).Standalone()

		bcmPricingCalculator_listBillEstimateInputUsageModificationsCmd.Flags().String("bill-estimate-id", "", "The unique identifier of the bill estimate to list input usage modifications for.")
		bcmPricingCalculator_listBillEstimateInputUsageModificationsCmd.Flags().String("filters", "", "Filters to apply to the list of input usage modifications.")
		bcmPricingCalculator_listBillEstimateInputUsageModificationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		bcmPricingCalculator_listBillEstimateInputUsageModificationsCmd.Flags().String("next-token", "", "A token to retrieve the next page of results.")
		bcmPricingCalculator_listBillEstimateInputUsageModificationsCmd.MarkFlagRequired("bill-estimate-id")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_listBillEstimateInputUsageModificationsCmd)
}
