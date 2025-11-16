package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_batchCreateWorkloadEstimateUsageCmd = &cobra.Command{
	Use:   "batch-create-workload-estimate-usage",
	Short: "Create Amazon Web Services service usage that you want to model in a Workload Estimate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_batchCreateWorkloadEstimateUsageCmd).Standalone()

	bcmPricingCalculator_batchCreateWorkloadEstimateUsageCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	bcmPricingCalculator_batchCreateWorkloadEstimateUsageCmd.Flags().String("usage", "", "List of usage that you want to model in the Workload estimate.")
	bcmPricingCalculator_batchCreateWorkloadEstimateUsageCmd.Flags().String("workload-estimate-id", "", "The ID of the Workload estimate for which you want to create the modeled usage.")
	bcmPricingCalculator_batchCreateWorkloadEstimateUsageCmd.MarkFlagRequired("usage")
	bcmPricingCalculator_batchCreateWorkloadEstimateUsageCmd.MarkFlagRequired("workload-estimate-id")
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_batchCreateWorkloadEstimateUsageCmd)
}
