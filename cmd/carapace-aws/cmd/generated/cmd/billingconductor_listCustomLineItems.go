package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_listCustomLineItemsCmd = &cobra.Command{
	Use:   "list-custom-line-items",
	Short: "A paginated call to get a list of all custom line items (FFLIs) for the given billing period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_listCustomLineItemsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_listCustomLineItemsCmd).Standalone()

		billingconductor_listCustomLineItemsCmd.Flags().String("billing-period", "", "The preferred billing period to get custom line items (FFLIs).")
		billingconductor_listCustomLineItemsCmd.Flags().String("filters", "", "A `ListCustomLineItemsFilter` that specifies the custom line item names and/or billing group Amazon Resource Names (ARNs) to retrieve FFLI information.")
		billingconductor_listCustomLineItemsCmd.Flags().String("max-results", "", "The maximum number of billing groups to retrieve.")
		billingconductor_listCustomLineItemsCmd.Flags().String("next-token", "", "The pagination token that's used on subsequent calls to get custom line items (FFLIs).")
	})
	billingconductorCmd.AddCommand(billingconductor_listCustomLineItemsCmd)
}
