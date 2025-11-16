package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_listWorkloadEstimatesCmd = &cobra.Command{
	Use:   "list-workload-estimates",
	Short: "Lists all workload estimates for the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_listWorkloadEstimatesCmd).Standalone()

	bcmPricingCalculator_listWorkloadEstimatesCmd.Flags().String("created-at-filter", "", "Filter workload estimates based on the creation date.")
	bcmPricingCalculator_listWorkloadEstimatesCmd.Flags().String("expires-at-filter", "", "Filter workload estimates based on the expiration date.")
	bcmPricingCalculator_listWorkloadEstimatesCmd.Flags().String("filters", "", "Filters to apply to the list of workload estimates.")
	bcmPricingCalculator_listWorkloadEstimatesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	bcmPricingCalculator_listWorkloadEstimatesCmd.Flags().String("next-token", "", "A token to retrieve the next page of results.")
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_listWorkloadEstimatesCmd)
}
