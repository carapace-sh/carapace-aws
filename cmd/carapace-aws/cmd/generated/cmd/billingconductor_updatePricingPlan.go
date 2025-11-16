package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_updatePricingPlanCmd = &cobra.Command{
	Use:   "update-pricing-plan",
	Short: "This updates an existing pricing plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_updatePricingPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_updatePricingPlanCmd).Standalone()

		billingconductor_updatePricingPlanCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the pricing plan that you're updating.")
		billingconductor_updatePricingPlanCmd.Flags().String("description", "", "The description of the pricing plan.")
		billingconductor_updatePricingPlanCmd.Flags().String("name", "", "The name of the pricing plan.")
		billingconductor_updatePricingPlanCmd.MarkFlagRequired("arn")
	})
	billingconductorCmd.AddCommand(billingconductor_updatePricingPlanCmd)
}
