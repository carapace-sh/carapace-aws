package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_updateWorkloadEstimateCmd = &cobra.Command{
	Use:   "update-workload-estimate",
	Short: "Updates an existing workload estimate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_updateWorkloadEstimateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_updateWorkloadEstimateCmd).Standalone()

		bcmPricingCalculator_updateWorkloadEstimateCmd.Flags().String("expires-at", "", "The new expiration date for the workload estimate.")
		bcmPricingCalculator_updateWorkloadEstimateCmd.Flags().String("identifier", "", "The unique identifier of the workload estimate to update.")
		bcmPricingCalculator_updateWorkloadEstimateCmd.Flags().String("name", "", "The new name for the workload estimate.")
		bcmPricingCalculator_updateWorkloadEstimateCmd.MarkFlagRequired("identifier")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_updateWorkloadEstimateCmd)
}
