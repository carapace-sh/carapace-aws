package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var invoicing_listInvoiceSummariesCmd = &cobra.Command{
	Use:   "list-invoice-summaries",
	Short: "Retrieves your invoice details programmatically, without line item details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoicing_listInvoiceSummariesCmd).Standalone()

	invoicing_listInvoiceSummariesCmd.Flags().String("filter", "", "Filters you can use to customize your invoice summary.")
	invoicing_listInvoiceSummariesCmd.Flags().String("max-results", "", "The maximum number of invoice summaries a paginated response can contain.")
	invoicing_listInvoiceSummariesCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
	invoicing_listInvoiceSummariesCmd.Flags().String("selector", "", "The option to retrieve details for a specific invoice by providing its unique ID.")
	invoicing_listInvoiceSummariesCmd.MarkFlagRequired("selector")
	invoicingCmd.AddCommand(invoicing_listInvoiceSummariesCmd)
}
