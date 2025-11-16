package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var invoicing_createInvoiceUnitCmd = &cobra.Command{
	Use:   "create-invoice-unit",
	Short: "This creates a new invoice unit with the provided definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoicing_createInvoiceUnitCmd).Standalone()

	invoicing_createInvoiceUnitCmd.Flags().String("description", "", "The invoice unit's description.")
	invoicing_createInvoiceUnitCmd.Flags().String("invoice-receiver", "", "The Amazon Web Services account ID chosen to be the receiver of an invoice unit.")
	invoicing_createInvoiceUnitCmd.Flags().String("name", "", "The unique name of the invoice unit that is shown on the generated invoice.")
	invoicing_createInvoiceUnitCmd.Flags().String("resource-tags", "", "The tag structure that contains a tag key and value.")
	invoicing_createInvoiceUnitCmd.Flags().String("rule", "", "The `InvoiceUnitRule` object used to create invoice units.")
	invoicing_createInvoiceUnitCmd.Flags().String("tax-inheritance-disabled", "", "Whether the invoice unit based tax inheritance is/ should be enabled or disabled.")
	invoicing_createInvoiceUnitCmd.MarkFlagRequired("invoice-receiver")
	invoicing_createInvoiceUnitCmd.MarkFlagRequired("name")
	invoicing_createInvoiceUnitCmd.MarkFlagRequired("rule")
	invoicingCmd.AddCommand(invoicing_createInvoiceUnitCmd)
}
