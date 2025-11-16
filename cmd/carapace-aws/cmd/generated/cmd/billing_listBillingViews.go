package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billing_listBillingViewsCmd = &cobra.Command{
	Use:   "list-billing-views",
	Short: "Lists the billing views available for a given time period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billing_listBillingViewsCmd).Standalone()

	billing_listBillingViewsCmd.Flags().String("active-time-range", "", "The time range for the billing views listed.")
	billing_listBillingViewsCmd.Flags().String("arns", "", "The Amazon Resource Name (ARN) that can be used to uniquely identify the billing view.")
	billing_listBillingViewsCmd.Flags().String("billing-view-types", "", "The type of billing view.")
	billing_listBillingViewsCmd.Flags().String("max-results", "", "The maximum number of billing views to retrieve.")
	billing_listBillingViewsCmd.Flags().String("next-token", "", "The pagination token that is used on subsequent calls to list billing views.")
	billing_listBillingViewsCmd.Flags().String("owner-account-id", "", "The list of owners of the billing view.")
	billing_listBillingViewsCmd.Flags().String("source-account-id", "", "Filters the results to include only billing views that use the specified account as a source.")
	billingCmd.AddCommand(billing_listBillingViewsCmd)
}
