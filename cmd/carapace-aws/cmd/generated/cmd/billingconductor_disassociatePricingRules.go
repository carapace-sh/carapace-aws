package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_disassociatePricingRulesCmd = &cobra.Command{
	Use:   "disassociate-pricing-rules",
	Short: "Disassociates a list of pricing rules from a pricing plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_disassociatePricingRulesCmd).Standalone()

	billingconductor_disassociatePricingRulesCmd.Flags().String("arn", "", "The pricing plan Amazon Resource Name (ARN) to disassociate pricing rules from.")
	billingconductor_disassociatePricingRulesCmd.Flags().String("pricing-rule-arns", "", "A list containing the Amazon Resource Name (ARN) of the pricing rules that will be disassociated.")
	billingconductor_disassociatePricingRulesCmd.MarkFlagRequired("arn")
	billingconductor_disassociatePricingRulesCmd.MarkFlagRequired("pricing-rule-arns")
	billingconductorCmd.AddCommand(billingconductor_disassociatePricingRulesCmd)
}
