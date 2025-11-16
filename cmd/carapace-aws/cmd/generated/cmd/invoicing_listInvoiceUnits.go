package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var invoicing_listInvoiceUnitsCmd = &cobra.Command{
	Use:   "list-invoice-units",
	Short: "This fetches a list of all invoice unit definitions for a given account, as of the provided `AsOf` date.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoicing_listInvoiceUnitsCmd).Standalone()

	invoicing_listInvoiceUnitsCmd.Flags().String("as-of", "", "The state of an invoice unit at a specified time.")
	invoicing_listInvoiceUnitsCmd.Flags().String("filters", "", "An optional input to the list API.")
	invoicing_listInvoiceUnitsCmd.Flags().String("max-results", "", "The maximum number of invoice units that can be returned.")
	invoicing_listInvoiceUnitsCmd.Flags().String("next-token", "", "The next token used to indicate where the returned list should start from.")
	invoicingCmd.AddCommand(invoicing_listInvoiceUnitsCmd)
}
