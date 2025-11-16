package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_batchDeleteWorkloadEstimateUsageCmd = &cobra.Command{
	Use:   "batch-delete-workload-estimate-usage",
	Short: "Delete usage that you have created in a Workload estimate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_batchDeleteWorkloadEstimateUsageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_batchDeleteWorkloadEstimateUsageCmd).Standalone()

		bcmPricingCalculator_batchDeleteWorkloadEstimateUsageCmd.Flags().String("ids", "", "List of usage that you want to delete from the Workload estimate.")
		bcmPricingCalculator_batchDeleteWorkloadEstimateUsageCmd.Flags().String("workload-estimate-id", "", "The ID of the Workload estimate for which you want to delete the modeled usage.")
		bcmPricingCalculator_batchDeleteWorkloadEstimateUsageCmd.MarkFlagRequired("ids")
		bcmPricingCalculator_batchDeleteWorkloadEstimateUsageCmd.MarkFlagRequired("workload-estimate-id")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_batchDeleteWorkloadEstimateUsageCmd)
}
