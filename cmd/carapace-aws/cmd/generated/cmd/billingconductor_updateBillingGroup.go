package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_updateBillingGroupCmd = &cobra.Command{
	Use:   "update-billing-group",
	Short: "This updates an existing billing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_updateBillingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_updateBillingGroupCmd).Standalone()

		billingconductor_updateBillingGroupCmd.Flags().String("account-grouping", "", "Specifies if the billing group has automatic account association (`AutoAssociate`) enabled.")
		billingconductor_updateBillingGroupCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the billing group being updated.")
		billingconductor_updateBillingGroupCmd.Flags().String("computation-preference", "", "The preferences and settings that will be used to compute the Amazon Web Services charges for a billing group.")
		billingconductor_updateBillingGroupCmd.Flags().String("description", "", "A description of the billing group.")
		billingconductor_updateBillingGroupCmd.Flags().String("name", "", "The name of the billing group.")
		billingconductor_updateBillingGroupCmd.Flags().String("status", "", "The status of the billing group.")
		billingconductor_updateBillingGroupCmd.MarkFlagRequired("arn")
	})
	billingconductorCmd.AddCommand(billingconductor_updateBillingGroupCmd)
}
