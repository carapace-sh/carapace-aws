package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_listWorkloadEstimateUsageCmd = &cobra.Command{
	Use:   "list-workload-estimate-usage",
	Short: "Lists the usage associated with a workload estimate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_listWorkloadEstimateUsageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_listWorkloadEstimateUsageCmd).Standalone()

		bcmPricingCalculator_listWorkloadEstimateUsageCmd.Flags().String("filters", "", "Filters to apply to the list of usage items.")
		bcmPricingCalculator_listWorkloadEstimateUsageCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		bcmPricingCalculator_listWorkloadEstimateUsageCmd.Flags().String("next-token", "", "A token to retrieve the next page of results.")
		bcmPricingCalculator_listWorkloadEstimateUsageCmd.Flags().String("workload-estimate-id", "", "The unique identifier of the workload estimate to list usage for.")
		bcmPricingCalculator_listWorkloadEstimateUsageCmd.MarkFlagRequired("workload-estimate-id")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_listWorkloadEstimateUsageCmd)
}
