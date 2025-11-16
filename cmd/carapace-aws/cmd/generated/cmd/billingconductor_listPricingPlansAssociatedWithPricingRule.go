package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_listPricingPlansAssociatedWithPricingRuleCmd = &cobra.Command{
	Use:   "list-pricing-plans-associated-with-pricing-rule",
	Short: "A list of the pricing plans that are associated with a pricing rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_listPricingPlansAssociatedWithPricingRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_listPricingPlansAssociatedWithPricingRuleCmd).Standalone()

		billingconductor_listPricingPlansAssociatedWithPricingRuleCmd.Flags().String("billing-period", "", "The pricing plan billing period for which associations will be listed.")
		billingconductor_listPricingPlansAssociatedWithPricingRuleCmd.Flags().String("max-results", "", "The optional maximum number of pricing rule associations to retrieve.")
		billingconductor_listPricingPlansAssociatedWithPricingRuleCmd.Flags().String("next-token", "", "The optional pagination token returned by a previous call.")
		billingconductor_listPricingPlansAssociatedWithPricingRuleCmd.Flags().String("pricing-rule-arn", "", "The pricing rule Amazon Resource Name (ARN) for which associations will be listed.")
		billingconductor_listPricingPlansAssociatedWithPricingRuleCmd.MarkFlagRequired("pricing-rule-arn")
	})
	billingconductorCmd.AddCommand(billingconductor_listPricingPlansAssociatedWithPricingRuleCmd)
}
