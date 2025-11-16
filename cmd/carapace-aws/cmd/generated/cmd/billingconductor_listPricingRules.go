package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_listPricingRulesCmd = &cobra.Command{
	Use:   "list-pricing-rules",
	Short: "Describes a pricing rule that can be associated to a pricing plan, or set of pricing plans.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_listPricingRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_listPricingRulesCmd).Standalone()

		billingconductor_listPricingRulesCmd.Flags().String("billing-period", "", "The preferred billing period to get the pricing plan.")
		billingconductor_listPricingRulesCmd.Flags().String("filters", "", "A `DescribePricingRuleFilter` that specifies the Amazon Resource Name (ARNs) of pricing rules to retrieve pricing rules information.")
		billingconductor_listPricingRulesCmd.Flags().String("max-results", "", "The maximum number of pricing rules to retrieve.")
		billingconductor_listPricingRulesCmd.Flags().String("next-token", "", "The pagination token that's used on subsequent call to get pricing rules.")
	})
	billingconductorCmd.AddCommand(billingconductor_listPricingRulesCmd)
}
