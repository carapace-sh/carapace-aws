package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_updatePricingRuleCmd = &cobra.Command{
	Use:   "update-pricing-rule",
	Short: "Updates an existing pricing rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_updatePricingRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_updatePricingRuleCmd).Standalone()

		billingconductor_updatePricingRuleCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the pricing rule to update.")
		billingconductor_updatePricingRuleCmd.Flags().String("description", "", "The new description for the pricing rule.")
		billingconductor_updatePricingRuleCmd.Flags().String("modifier-percentage", "", "The new modifier to show pricing plan rates as a percentage.")
		billingconductor_updatePricingRuleCmd.Flags().String("name", "", "The new name of the pricing rule.")
		billingconductor_updatePricingRuleCmd.Flags().String("tiering", "", "The set of tiering configurations for the pricing rule.")
		billingconductor_updatePricingRuleCmd.Flags().String("type", "", "The new pricing rule type.")
		billingconductor_updatePricingRuleCmd.MarkFlagRequired("arn")
	})
	billingconductorCmd.AddCommand(billingconductor_updatePricingRuleCmd)
}
