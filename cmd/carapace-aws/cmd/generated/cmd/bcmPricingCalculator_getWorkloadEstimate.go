package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_getWorkloadEstimateCmd = &cobra.Command{
	Use:   "get-workload-estimate",
	Short: "Retrieves details of a specific workload estimate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_getWorkloadEstimateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_getWorkloadEstimateCmd).Standalone()

		bcmPricingCalculator_getWorkloadEstimateCmd.Flags().String("identifier", "", "The unique identifier of the workload estimate to retrieve.")
		bcmPricingCalculator_getWorkloadEstimateCmd.MarkFlagRequired("identifier")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_getWorkloadEstimateCmd)
}
