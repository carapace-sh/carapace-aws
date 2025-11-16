package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_listTransactionsCmd = &cobra.Command{
	Use:   "list-transactions",
	Short: "Returns metadata about transactions and their status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_listTransactionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_listTransactionsCmd).Standalone()

		lakeformation_listTransactionsCmd.Flags().String("catalog-id", "", "The catalog for which to list transactions.")
		lakeformation_listTransactionsCmd.Flags().String("max-results", "", "The maximum number of transactions to return in a single call.")
		lakeformation_listTransactionsCmd.Flags().String("next-token", "", "A continuation token if this is not the first call to retrieve transactions.")
		lakeformation_listTransactionsCmd.Flags().String("status-filter", "", "A filter indicating the status of transactions to return.")
	})
	lakeformationCmd.AddCommand(lakeformation_listTransactionsCmd)
}
