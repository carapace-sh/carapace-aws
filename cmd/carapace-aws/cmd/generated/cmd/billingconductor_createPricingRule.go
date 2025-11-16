package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_createPricingRuleCmd = &cobra.Command{
	Use:   "create-pricing-rule",
	Short: "Creates a pricing rule can be associated to a pricing plan, or a set of pricing plans.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_createPricingRuleCmd).Standalone()

	billingconductor_createPricingRuleCmd.Flags().String("billing-entity", "", "The seller of services provided by Amazon Web Services, their affiliates, or third-party providers selling services via Amazon Web Services Marketplace.")
	billingconductor_createPricingRuleCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you specify to ensure idempotency of the request.")
	billingconductor_createPricingRuleCmd.Flags().String("description", "", "The pricing rule description.")
	billingconductor_createPricingRuleCmd.Flags().String("modifier-percentage", "", "A percentage modifier that's applied on the public pricing rates.")
	billingconductor_createPricingRuleCmd.Flags().String("name", "", "The pricing rule name.")
	billingconductor_createPricingRuleCmd.Flags().String("operation", "", "Operation is the specific Amazon Web Services action covered by this line item.")
	billingconductor_createPricingRuleCmd.Flags().String("scope", "", "The scope of pricing rule that indicates if it's globally applicable, or it's service-specific.")
	billingconductor_createPricingRuleCmd.Flags().String("service", "", "If the `Scope` attribute is set to `SERVICE` or `SKU`, the attribute indicates which service the `PricingRule` is applicable for.")
	billingconductor_createPricingRuleCmd.Flags().String("tags", "", "A map that contains tag keys and tag values that are attached to a pricing rule.")
	billingconductor_createPricingRuleCmd.Flags().String("tiering", "", "The set of tiering configurations for the pricing rule.")
	billingconductor_createPricingRuleCmd.Flags().String("type", "", "The type of pricing rule.")
	billingconductor_createPricingRuleCmd.Flags().String("usage-type", "", "Usage type is the unit that each service uses to measure the usage of a specific type of resource.")
	billingconductor_createPricingRuleCmd.MarkFlagRequired("name")
	billingconductor_createPricingRuleCmd.MarkFlagRequired("scope")
	billingconductor_createPricingRuleCmd.MarkFlagRequired("type")
	billingconductorCmd.AddCommand(billingconductor_createPricingRuleCmd)
}
