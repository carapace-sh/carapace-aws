package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_deleteWorkloadEstimateCmd = &cobra.Command{
	Use:   "delete-workload-estimate",
	Short: "Deletes an existing workload estimate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_deleteWorkloadEstimateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_deleteWorkloadEstimateCmd).Standalone()

		bcmPricingCalculator_deleteWorkloadEstimateCmd.Flags().String("identifier", "", "The unique identifier of the workload estimate to delete.")
		bcmPricingCalculator_deleteWorkloadEstimateCmd.MarkFlagRequired("identifier")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_deleteWorkloadEstimateCmd)
}
