package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_listPricingPlansCmd = &cobra.Command{
	Use:   "list-pricing-plans",
	Short: "A paginated call to get pricing plans for the given billing period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_listPricingPlansCmd).Standalone()

	billingconductor_listPricingPlansCmd.Flags().String("billing-period", "", "The preferred billing period to get pricing plan.")
	billingconductor_listPricingPlansCmd.Flags().String("filters", "", "A `ListPricingPlansFilter` that specifies the Amazon Resource Name (ARNs) of pricing plans to retrieve pricing plans information.")
	billingconductor_listPricingPlansCmd.Flags().String("max-results", "", "The maximum number of pricing plans to retrieve.")
	billingconductor_listPricingPlansCmd.Flags().String("next-token", "", "The pagination token that's used on subsequent call to get pricing plans.")
	billingconductorCmd.AddCommand(billingconductor_listPricingPlansCmd)
}
