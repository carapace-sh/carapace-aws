package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var invoicing_getInvoiceUnitCmd = &cobra.Command{
	Use:   "get-invoice-unit",
	Short: "This retrieves the invoice unit definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoicing_getInvoiceUnitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(invoicing_getInvoiceUnitCmd).Standalone()

		invoicing_getInvoiceUnitCmd.Flags().String("as-of", "", "The state of an invoice unit at a specified time.")
		invoicing_getInvoiceUnitCmd.Flags().String("invoice-unit-arn", "", "The ARN to identify an invoice unit.")
		invoicing_getInvoiceUnitCmd.MarkFlagRequired("invoice-unit-arn")
	})
	invoicingCmd.AddCommand(invoicing_getInvoiceUnitCmd)
}
