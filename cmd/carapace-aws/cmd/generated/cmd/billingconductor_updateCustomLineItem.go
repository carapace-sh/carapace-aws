package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_updateCustomLineItemCmd = &cobra.Command{
	Use:   "update-custom-line-item",
	Short: "Update an existing custom line item in the current or previous billing period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_updateCustomLineItemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_updateCustomLineItemCmd).Standalone()

		billingconductor_updateCustomLineItemCmd.Flags().String("arn", "", "The ARN of the custom line item to be updated.")
		billingconductor_updateCustomLineItemCmd.Flags().String("billing-period-range", "", "")
		billingconductor_updateCustomLineItemCmd.Flags().String("charge-details", "", "A `ListCustomLineItemChargeDetails` containing the new charge details for the custom line item.")
		billingconductor_updateCustomLineItemCmd.Flags().String("description", "", "The new line item description of the custom line item.")
		billingconductor_updateCustomLineItemCmd.Flags().String("name", "", "The new name for the custom line item.")
		billingconductor_updateCustomLineItemCmd.MarkFlagRequired("arn")
	})
	billingconductorCmd.AddCommand(billingconductor_updateCustomLineItemCmd)
}
