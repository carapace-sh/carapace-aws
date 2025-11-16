package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_deleteCustomLineItemCmd = &cobra.Command{
	Use:   "delete-custom-line-item",
	Short: "Deletes the custom line item identified by the given ARN in the current, or previous billing period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_deleteCustomLineItemCmd).Standalone()

	billingconductor_deleteCustomLineItemCmd.Flags().String("arn", "", "The ARN of the custom line item to be deleted.")
	billingconductor_deleteCustomLineItemCmd.Flags().String("billing-period-range", "", "")
	billingconductor_deleteCustomLineItemCmd.MarkFlagRequired("arn")
	billingconductorCmd.AddCommand(billingconductor_deleteCustomLineItemCmd)
}
