package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var invoicing_getInvoicePdfCmd = &cobra.Command{
	Use:   "get-invoice-pdf",
	Short: "Returns a URL to download the invoice document and supplemental documents associated with an invoice.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoicing_getInvoicePdfCmd).Standalone()

	invoicing_getInvoicePdfCmd.Flags().String("invoice-id", "", "Your unique invoice ID.")
	invoicing_getInvoicePdfCmd.MarkFlagRequired("invoice-id")
	invoicingCmd.AddCommand(invoicing_getInvoicePdfCmd)
}
