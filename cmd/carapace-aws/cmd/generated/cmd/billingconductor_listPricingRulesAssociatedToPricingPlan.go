package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_listPricingRulesAssociatedToPricingPlanCmd = &cobra.Command{
	Use:   "list-pricing-rules-associated-to-pricing-plan",
	Short: "Lists the pricing rules that are associated with a pricing plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_listPricingRulesAssociatedToPricingPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_listPricingRulesAssociatedToPricingPlanCmd).Standalone()

		billingconductor_listPricingRulesAssociatedToPricingPlanCmd.Flags().String("billing-period", "", "The billing period for which the pricing rule associations are to be listed.")
		billingconductor_listPricingRulesAssociatedToPricingPlanCmd.Flags().String("max-results", "", "The optional maximum number of pricing rule associations to retrieve.")
		billingconductor_listPricingRulesAssociatedToPricingPlanCmd.Flags().String("next-token", "", "The optional pagination token returned by a previous call.")
		billingconductor_listPricingRulesAssociatedToPricingPlanCmd.Flags().String("pricing-plan-arn", "", "The Amazon Resource Name (ARN) of the pricing plan for which associations are to be listed.")
		billingconductor_listPricingRulesAssociatedToPricingPlanCmd.MarkFlagRequired("pricing-plan-arn")
	})
	billingconductorCmd.AddCommand(billingconductor_listPricingRulesAssociatedToPricingPlanCmd)
}
