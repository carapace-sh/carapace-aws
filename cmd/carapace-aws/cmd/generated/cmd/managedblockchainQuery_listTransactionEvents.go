package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchainQuery_listTransactionEventsCmd = &cobra.Command{
	Use:   "list-transaction-events",
	Short: "Lists all the transaction events for a transaction\n\nThis action will return transaction details for all transactions that are *confirmed* on the blockchain, even if they have not reached [finality](https://docs.aws.amazon.com/managed-blockchain/latest/ambq-dg/key-concepts.html#finality).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchainQuery_listTransactionEventsCmd).Standalone()

	managedblockchainQuery_listTransactionEventsCmd.Flags().String("max-results", "", "The maximum number of transaction events to list.")
	managedblockchainQuery_listTransactionEventsCmd.Flags().String("network", "", "The blockchain network where the transaction events occurred.")
	managedblockchainQuery_listTransactionEventsCmd.Flags().String("next-token", "", "The pagination token that indicates the next set of results to retrieve.")
	managedblockchainQuery_listTransactionEventsCmd.Flags().String("transaction-hash", "", "The hash of a transaction.")
	managedblockchainQuery_listTransactionEventsCmd.Flags().String("transaction-id", "", "The identifier of a Bitcoin transaction.")
	managedblockchainQuery_listTransactionEventsCmd.MarkFlagRequired("network")
	managedblockchainQueryCmd.AddCommand(managedblockchainQuery_listTransactionEventsCmd)
}
