package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_associatePricingRulesCmd = &cobra.Command{
	Use:   "associate-pricing-rules",
	Short: "Connects an array of `PricingRuleArns` to a defined `PricingPlan`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_associatePricingRulesCmd).Standalone()

	billingconductor_associatePricingRulesCmd.Flags().String("arn", "", "The `PricingPlanArn` that the `PricingRuleArns` are associated with.")
	billingconductor_associatePricingRulesCmd.Flags().String("pricing-rule-arns", "", "The `PricingRuleArns` that are associated with the Pricing Plan.")
	billingconductor_associatePricingRulesCmd.MarkFlagRequired("arn")
	billingconductor_associatePricingRulesCmd.MarkFlagRequired("pricing-rule-arns")
	billingconductorCmd.AddCommand(billingconductor_associatePricingRulesCmd)
}
