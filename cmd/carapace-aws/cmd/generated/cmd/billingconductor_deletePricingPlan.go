package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_deletePricingPlanCmd = &cobra.Command{
	Use:   "delete-pricing-plan",
	Short: "Deletes a pricing plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_deletePricingPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_deletePricingPlanCmd).Standalone()

		billingconductor_deletePricingPlanCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the pricing plan that you're deleting.")
		billingconductor_deletePricingPlanCmd.MarkFlagRequired("arn")
	})
	billingconductorCmd.AddCommand(billingconductor_deletePricingPlanCmd)
}
