package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billing_updateBillingViewCmd = &cobra.Command{
	Use:   "update-billing-view",
	Short: "An API to update the attributes of the billing view.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billing_updateBillingViewCmd).Standalone()

	billing_updateBillingViewCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) that can be used to uniquely identify the billing view.")
	billing_updateBillingViewCmd.Flags().String("data-filter-expression", "", "See [Expression](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_billing_Expression.html).")
	billing_updateBillingViewCmd.Flags().String("description", "", "The description of the billing view.")
	billing_updateBillingViewCmd.Flags().String("name", "", "The name of the billing view.")
	billing_updateBillingViewCmd.MarkFlagRequired("arn")
	billingCmd.AddCommand(billing_updateBillingViewCmd)
}
