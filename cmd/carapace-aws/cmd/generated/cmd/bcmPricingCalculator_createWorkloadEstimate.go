package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_createWorkloadEstimateCmd = &cobra.Command{
	Use:   "create-workload-estimate",
	Short: "Creates a new workload estimate to model costs for a specific workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_createWorkloadEstimateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_createWorkloadEstimateCmd).Standalone()

		bcmPricingCalculator_createWorkloadEstimateCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure idempotency of the request.")
		bcmPricingCalculator_createWorkloadEstimateCmd.Flags().String("name", "", "A descriptive name for the workload estimate.")
		bcmPricingCalculator_createWorkloadEstimateCmd.Flags().String("rate-type", "", "The type of pricing rates to use for the estimate.")
		bcmPricingCalculator_createWorkloadEstimateCmd.Flags().String("tags", "", "The tags to apply to the workload estimate.")
		bcmPricingCalculator_createWorkloadEstimateCmd.MarkFlagRequired("name")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_createWorkloadEstimateCmd)
}
