package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var savingsplans_describeSavingsPlansOfferingsCmd = &cobra.Command{
	Use:   "describe-savings-plans-offerings",
	Short: "Describes the offerings for the specified Savings Plans.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(savingsplans_describeSavingsPlansOfferingsCmd).Standalone()

	savingsplans_describeSavingsPlansOfferingsCmd.Flags().String("currencies", "", "The currencies.")
	savingsplans_describeSavingsPlansOfferingsCmd.Flags().String("descriptions", "", "The descriptions.")
	savingsplans_describeSavingsPlansOfferingsCmd.Flags().String("durations", "", "The duration, in seconds.")
	savingsplans_describeSavingsPlansOfferingsCmd.Flags().String("filters", "", "The filters.")
	savingsplans_describeSavingsPlansOfferingsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	savingsplans_describeSavingsPlansOfferingsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	savingsplans_describeSavingsPlansOfferingsCmd.Flags().String("offering-ids", "", "The IDs of the offerings.")
	savingsplans_describeSavingsPlansOfferingsCmd.Flags().String("operations", "", "The specific Amazon Web Services operation for the line item in the billing report.")
	savingsplans_describeSavingsPlansOfferingsCmd.Flags().String("payment-options", "", "The payment options.")
	savingsplans_describeSavingsPlansOfferingsCmd.Flags().String("plan-types", "", "The plan types.")
	savingsplans_describeSavingsPlansOfferingsCmd.Flags().String("product-type", "", "The product type.")
	savingsplans_describeSavingsPlansOfferingsCmd.Flags().String("service-codes", "", "The services.")
	savingsplans_describeSavingsPlansOfferingsCmd.Flags().String("usage-types", "", "The usage details of the line item in the billing report.")
	savingsplansCmd.AddCommand(savingsplans_describeSavingsPlansOfferingsCmd)
}
