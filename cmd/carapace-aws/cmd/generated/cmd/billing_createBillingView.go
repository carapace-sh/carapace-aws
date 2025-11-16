package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billing_createBillingViewCmd = &cobra.Command{
	Use:   "create-billing-view",
	Short: "Creates a billing view with the specified billing view attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billing_createBillingViewCmd).Standalone()

	billing_createBillingViewCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier you specify to ensure idempotency of the request.")
	billing_createBillingViewCmd.Flags().String("data-filter-expression", "", "See [Expression](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_billing_Expression.html).")
	billing_createBillingViewCmd.Flags().String("description", "", "The description of the billing view.")
	billing_createBillingViewCmd.Flags().String("name", "", "The name of the billing view.")
	billing_createBillingViewCmd.Flags().String("resource-tags", "", "A list of key value map specifying tags associated to the billing view being created.")
	billing_createBillingViewCmd.Flags().String("source-views", "", "A list of billing views used as the data source for the custom billing view.")
	billing_createBillingViewCmd.MarkFlagRequired("name")
	billing_createBillingViewCmd.MarkFlagRequired("source-views")
	billingCmd.AddCommand(billing_createBillingViewCmd)
}
