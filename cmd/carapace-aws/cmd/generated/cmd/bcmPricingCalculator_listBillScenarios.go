package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_listBillScenariosCmd = &cobra.Command{
	Use:   "list-bill-scenarios",
	Short: "Lists all bill scenarios for the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_listBillScenariosCmd).Standalone()

	bcmPricingCalculator_listBillScenariosCmd.Flags().String("created-at-filter", "", "Filter bill scenarios based on the creation date.")
	bcmPricingCalculator_listBillScenariosCmd.Flags().String("expires-at-filter", "", "Filter bill scenarios based on the expiration date.")
	bcmPricingCalculator_listBillScenariosCmd.Flags().String("filters", "", "Filters to apply to the list of bill scenarios.")
	bcmPricingCalculator_listBillScenariosCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	bcmPricingCalculator_listBillScenariosCmd.Flags().String("next-token", "", "A token to retrieve the next page of results.")
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_listBillScenariosCmd)
}
