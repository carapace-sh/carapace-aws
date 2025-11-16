package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_listBillScenarioUsageModificationsCmd = &cobra.Command{
	Use:   "list-bill-scenario-usage-modifications",
	Short: "Lists the usage modifications associated with a bill scenario.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_listBillScenarioUsageModificationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_listBillScenarioUsageModificationsCmd).Standalone()

		bcmPricingCalculator_listBillScenarioUsageModificationsCmd.Flags().String("bill-scenario-id", "", "The unique identifier of the bill scenario to list usage modifications for.")
		bcmPricingCalculator_listBillScenarioUsageModificationsCmd.Flags().String("filters", "", "Filters to apply to the list of usage modifications.")
		bcmPricingCalculator_listBillScenarioUsageModificationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		bcmPricingCalculator_listBillScenarioUsageModificationsCmd.Flags().String("next-token", "", "A token to retrieve the next page of results.")
		bcmPricingCalculator_listBillScenarioUsageModificationsCmd.MarkFlagRequired("bill-scenario-id")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_listBillScenarioUsageModificationsCmd)
}
