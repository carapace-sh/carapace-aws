package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var invoicing_batchGetInvoiceProfileCmd = &cobra.Command{
	Use:   "batch-get-invoice-profile",
	Short: "This gets the invoice profile associated with a set of accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoicing_batchGetInvoiceProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(invoicing_batchGetInvoiceProfileCmd).Standalone()

		invoicing_batchGetInvoiceProfileCmd.Flags().String("account-ids", "", "Retrieves the corresponding invoice profile data for these account IDs.")
		invoicing_batchGetInvoiceProfileCmd.MarkFlagRequired("account-ids")
	})
	invoicingCmd.AddCommand(invoicing_batchGetInvoiceProfileCmd)
}
