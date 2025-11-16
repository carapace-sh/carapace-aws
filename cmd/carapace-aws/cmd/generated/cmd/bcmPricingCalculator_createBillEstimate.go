package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_createBillEstimateCmd = &cobra.Command{
	Use:   "create-bill-estimate",
	Short: "Create a Bill estimate from a Bill scenario.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_createBillEstimateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_createBillEstimateCmd).Standalone()

		bcmPricingCalculator_createBillEstimateCmd.Flags().String("bill-scenario-id", "", "The ID of the Bill Scenario for which you want to create a Bill estimate.")
		bcmPricingCalculator_createBillEstimateCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		bcmPricingCalculator_createBillEstimateCmd.Flags().String("name", "", "The name of the Bill estimate that will be created.")
		bcmPricingCalculator_createBillEstimateCmd.Flags().String("tags", "", "An optional list of tags to associate with the specified BillEstimate.")
		bcmPricingCalculator_createBillEstimateCmd.MarkFlagRequired("bill-scenario-id")
		bcmPricingCalculator_createBillEstimateCmd.MarkFlagRequired("name")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_createBillEstimateCmd)
}
