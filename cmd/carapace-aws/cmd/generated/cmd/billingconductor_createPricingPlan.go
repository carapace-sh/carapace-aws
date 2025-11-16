package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_createPricingPlanCmd = &cobra.Command{
	Use:   "create-pricing-plan",
	Short: "Creates a pricing plan that is used for computing Amazon Web Services charges for billing groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_createPricingPlanCmd).Standalone()

	billingconductor_createPricingPlanCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you specify to ensure idempotency of the request.")
	billingconductor_createPricingPlanCmd.Flags().String("description", "", "The description of the pricing plan.")
	billingconductor_createPricingPlanCmd.Flags().String("name", "", "The name of the pricing plan.")
	billingconductor_createPricingPlanCmd.Flags().String("pricing-rule-arns", "", "A list of Amazon Resource Names (ARNs) that define the pricing plan parameters.")
	billingconductor_createPricingPlanCmd.Flags().String("tags", "", "A map that contains tag keys and tag values that are attached to a pricing plan.")
	billingconductor_createPricingPlanCmd.MarkFlagRequired("name")
	billingconductorCmd.AddCommand(billingconductor_createPricingPlanCmd)
}
