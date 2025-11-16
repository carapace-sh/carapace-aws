package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_createBillingGroupCmd = &cobra.Command{
	Use:   "create-billing-group",
	Short: "Creates a billing group that resembles a consolidated billing family that Amazon Web Services charges, based off of the predefined pricing plan computation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_createBillingGroupCmd).Standalone()

	billingconductor_createBillingGroupCmd.Flags().String("account-grouping", "", "The set of accounts that will be under the billing group.")
	billingconductor_createBillingGroupCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you specify to ensure idempotency of the request.")
	billingconductor_createBillingGroupCmd.Flags().String("computation-preference", "", "The preferences and settings that will be used to compute the Amazon Web Services charges for a billing group.")
	billingconductor_createBillingGroupCmd.Flags().String("description", "", "The description of the billing group.")
	billingconductor_createBillingGroupCmd.Flags().String("name", "", "The billing group name.")
	billingconductor_createBillingGroupCmd.Flags().String("primary-account-id", "", "The account ID that serves as the main account in a billing group.")
	billingconductor_createBillingGroupCmd.Flags().String("tags", "", "A map that contains tag keys and tag values that are attached to a billing group.")
	billingconductor_createBillingGroupCmd.MarkFlagRequired("account-grouping")
	billingconductor_createBillingGroupCmd.MarkFlagRequired("computation-preference")
	billingconductor_createBillingGroupCmd.MarkFlagRequired("name")
	billingconductorCmd.AddCommand(billingconductor_createBillingGroupCmd)
}
