package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var invoicing_deleteInvoiceUnitCmd = &cobra.Command{
	Use:   "delete-invoice-unit",
	Short: "This deletes an invoice unit with the provided invoice unit ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoicing_deleteInvoiceUnitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(invoicing_deleteInvoiceUnitCmd).Standalone()

		invoicing_deleteInvoiceUnitCmd.Flags().String("invoice-unit-arn", "", "The ARN to identify an invoice unit.")
		invoicing_deleteInvoiceUnitCmd.MarkFlagRequired("invoice-unit-arn")
	})
	invoicingCmd.AddCommand(invoicing_deleteInvoiceUnitCmd)
}
