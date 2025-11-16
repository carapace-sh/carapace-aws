package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_createCustomLineItemCmd = &cobra.Command{
	Use:   "create-custom-line-item",
	Short: "Creates a custom line item that can be used to create a one-time fixed charge that can be applied to a single billing group for the current or previous billing period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_createCustomLineItemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_createCustomLineItemCmd).Standalone()

		billingconductor_createCustomLineItemCmd.Flags().String("account-id", "", "The Amazon Web Services account in which this custom line item will be applied to.")
		billingconductor_createCustomLineItemCmd.Flags().String("billing-group-arn", "", "The Amazon Resource Name (ARN) that references the billing group where the custom line item applies to.")
		billingconductor_createCustomLineItemCmd.Flags().String("billing-period-range", "", "A time range for which the custom line item is effective.")
		billingconductor_createCustomLineItemCmd.Flags().String("charge-details", "", "A `CustomLineItemChargeDetails` that describes the charge details for a custom line item.")
		billingconductor_createCustomLineItemCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you specify to ensure idempotency of the request.")
		billingconductor_createCustomLineItemCmd.Flags().String("computation-rule", "", "")
		billingconductor_createCustomLineItemCmd.Flags().String("description", "", "The description of the custom line item.")
		billingconductor_createCustomLineItemCmd.Flags().String("name", "", "The name of the custom line item.")
		billingconductor_createCustomLineItemCmd.Flags().String("presentation-details", "", "")
		billingconductor_createCustomLineItemCmd.Flags().String("tags", "", "A map that contains tag keys and tag values that are attached to a custom line item.")
		billingconductor_createCustomLineItemCmd.MarkFlagRequired("billing-group-arn")
		billingconductor_createCustomLineItemCmd.MarkFlagRequired("charge-details")
		billingconductor_createCustomLineItemCmd.MarkFlagRequired("description")
		billingconductor_createCustomLineItemCmd.MarkFlagRequired("name")
	})
	billingconductorCmd.AddCommand(billingconductor_createCustomLineItemCmd)
}
