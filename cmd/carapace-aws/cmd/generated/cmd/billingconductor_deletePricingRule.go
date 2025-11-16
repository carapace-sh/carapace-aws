package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_deletePricingRuleCmd = &cobra.Command{
	Use:   "delete-pricing-rule",
	Short: "Deletes the pricing rule that's identified by the input Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_deletePricingRuleCmd).Standalone()

	billingconductor_deletePricingRuleCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the pricing rule that you are deleting.")
	billingconductor_deletePricingRuleCmd.MarkFlagRequired("arn")
	billingconductorCmd.AddCommand(billingconductor_deletePricingRuleCmd)
}
