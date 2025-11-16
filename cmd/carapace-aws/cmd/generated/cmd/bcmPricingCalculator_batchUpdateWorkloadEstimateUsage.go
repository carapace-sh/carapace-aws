package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_batchUpdateWorkloadEstimateUsageCmd = &cobra.Command{
	Use:   "batch-update-workload-estimate-usage",
	Short: "Update a newly added or existing usage lines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_batchUpdateWorkloadEstimateUsageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_batchUpdateWorkloadEstimateUsageCmd).Standalone()

		bcmPricingCalculator_batchUpdateWorkloadEstimateUsageCmd.Flags().String("usage", "", "List of usage line amounts and usage group that you want to update in a Workload estimate identified by the usage ID.")
		bcmPricingCalculator_batchUpdateWorkloadEstimateUsageCmd.Flags().String("workload-estimate-id", "", "The ID of the Workload estimate for which you want to modify the usage lines.")
		bcmPricingCalculator_batchUpdateWorkloadEstimateUsageCmd.MarkFlagRequired("usage")
		bcmPricingCalculator_batchUpdateWorkloadEstimateUsageCmd.MarkFlagRequired("workload-estimate-id")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_batchUpdateWorkloadEstimateUsageCmd)
}
