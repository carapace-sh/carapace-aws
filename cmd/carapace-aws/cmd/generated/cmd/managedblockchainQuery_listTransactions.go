package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchainQuery_listTransactionsCmd = &cobra.Command{
	Use:   "list-transactions",
	Short: "Lists all the transaction events for a transaction.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchainQuery_listTransactionsCmd).Standalone()

	managedblockchainQuery_listTransactionsCmd.Flags().String("address", "", "The address (either a contract or wallet), whose transactions are being requested.")
	managedblockchainQuery_listTransactionsCmd.Flags().String("confirmation-status-filter", "", "This filter is used to include transactions in the response that haven't reached [*finality*](https://docs.aws.amazon.com/managed-blockchain/latest/ambq-dg/key-concepts.html#finality) .")
	managedblockchainQuery_listTransactionsCmd.Flags().String("from-blockchain-instant", "", "")
	managedblockchainQuery_listTransactionsCmd.Flags().String("max-results", "", "The maximum number of transactions to list.")
	managedblockchainQuery_listTransactionsCmd.Flags().String("network", "", "The blockchain network where the transactions occurred.")
	managedblockchainQuery_listTransactionsCmd.Flags().String("next-token", "", "The pagination token that indicates the next set of results to retrieve.")
	managedblockchainQuery_listTransactionsCmd.Flags().String("sort", "", "The order by which the results will be sorted.")
	managedblockchainQuery_listTransactionsCmd.Flags().String("to-blockchain-instant", "", "")
	managedblockchainQuery_listTransactionsCmd.MarkFlagRequired("address")
	managedblockchainQuery_listTransactionsCmd.MarkFlagRequired("network")
	managedblockchainQueryCmd.AddCommand(managedblockchainQuery_listTransactionsCmd)
}
