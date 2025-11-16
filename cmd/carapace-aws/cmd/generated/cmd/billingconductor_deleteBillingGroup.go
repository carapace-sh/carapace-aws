package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_deleteBillingGroupCmd = &cobra.Command{
	Use:   "delete-billing-group",
	Short: "Deletes a billing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_deleteBillingGroupCmd).Standalone()

	billingconductor_deleteBillingGroupCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the billing group that you're deleting.")
	billingconductor_deleteBillingGroupCmd.MarkFlagRequired("arn")
	billingconductorCmd.AddCommand(billingconductor_deleteBillingGroupCmd)
}
