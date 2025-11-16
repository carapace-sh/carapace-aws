package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var savingsplans_describeSavingsPlansOfferingRatesCmd = &cobra.Command{
	Use:   "describe-savings-plans-offering-rates",
	Short: "Describes the offering rates for the specified Savings Plans.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(savingsplans_describeSavingsPlansOfferingRatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(savingsplans_describeSavingsPlansOfferingRatesCmd).Standalone()

		savingsplans_describeSavingsPlansOfferingRatesCmd.Flags().String("filters", "", "The filters.")
		savingsplans_describeSavingsPlansOfferingRatesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		savingsplans_describeSavingsPlansOfferingRatesCmd.Flags().String("next-token", "", "The token for the next page of results.")
		savingsplans_describeSavingsPlansOfferingRatesCmd.Flags().String("operations", "", "The specific Amazon Web Services operation for the line item in the billing report.")
		savingsplans_describeSavingsPlansOfferingRatesCmd.Flags().String("products", "", "The Amazon Web Services products.")
		savingsplans_describeSavingsPlansOfferingRatesCmd.Flags().String("savings-plan-offering-ids", "", "The IDs of the offerings.")
		savingsplans_describeSavingsPlansOfferingRatesCmd.Flags().String("savings-plan-payment-options", "", "The payment options.")
		savingsplans_describeSavingsPlansOfferingRatesCmd.Flags().String("savings-plan-types", "", "The plan types.")
		savingsplans_describeSavingsPlansOfferingRatesCmd.Flags().String("service-codes", "", "The services.")
		savingsplans_describeSavingsPlansOfferingRatesCmd.Flags().String("usage-types", "", "The usage details of the line item in the billing report.")
	})
	savingsplansCmd.AddCommand(savingsplans_describeSavingsPlansOfferingRatesCmd)
}
