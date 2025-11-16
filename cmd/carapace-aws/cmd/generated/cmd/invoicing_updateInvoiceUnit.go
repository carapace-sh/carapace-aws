package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var invoicing_updateInvoiceUnitCmd = &cobra.Command{
	Use:   "update-invoice-unit",
	Short: "You can update the invoice unit configuration at any time, and Amazon Web Services will use the latest configuration at the end of the month.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoicing_updateInvoiceUnitCmd).Standalone()

	invoicing_updateInvoiceUnitCmd.Flags().String("description", "", "The assigned description for an invoice unit.")
	invoicing_updateInvoiceUnitCmd.Flags().String("invoice-unit-arn", "", "The ARN to identify an invoice unit.")
	invoicing_updateInvoiceUnitCmd.Flags().String("rule", "", "The `InvoiceUnitRule` object used to update invoice units.")
	invoicing_updateInvoiceUnitCmd.Flags().String("tax-inheritance-disabled", "", "Whether the invoice unit based tax inheritance is/ should be enabled or disabled.")
	invoicing_updateInvoiceUnitCmd.MarkFlagRequired("invoice-unit-arn")
	invoicingCmd.AddCommand(invoicing_updateInvoiceUnitCmd)
}
