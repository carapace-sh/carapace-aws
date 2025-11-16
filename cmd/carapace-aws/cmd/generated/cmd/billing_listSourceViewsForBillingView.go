package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billing_listSourceViewsForBillingViewCmd = &cobra.Command{
	Use:   "list-source-views-for-billing-view",
	Short: "Lists the source views (managed Amazon Web Services billing views) associated with the billing view.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billing_listSourceViewsForBillingViewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billing_listSourceViewsForBillingViewCmd).Standalone()

		billing_listSourceViewsForBillingViewCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) that can be used to uniquely identify the billing view.")
		billing_listSourceViewsForBillingViewCmd.Flags().String("max-results", "", "The number of entries a paginated response contains.")
		billing_listSourceViewsForBillingViewCmd.Flags().String("next-token", "", "The pagination token that is used on subsequent calls to list billing views.")
		billing_listSourceViewsForBillingViewCmd.MarkFlagRequired("arn")
	})
	billingCmd.AddCommand(billing_listSourceViewsForBillingViewCmd)
}
