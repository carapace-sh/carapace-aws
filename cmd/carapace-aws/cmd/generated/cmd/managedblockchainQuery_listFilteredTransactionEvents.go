package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchainQuery_listFilteredTransactionEventsCmd = &cobra.Command{
	Use:   "list-filtered-transaction-events",
	Short: "Lists all the transaction events for an address on the blockchain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchainQuery_listFilteredTransactionEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(managedblockchainQuery_listFilteredTransactionEventsCmd).Standalone()

		managedblockchainQuery_listFilteredTransactionEventsCmd.Flags().String("address-identifier-filter", "", "This is the unique public address on the blockchain for which the transaction events are being requested.")
		managedblockchainQuery_listFilteredTransactionEventsCmd.Flags().String("confirmation-status-filter", "", "")
		managedblockchainQuery_listFilteredTransactionEventsCmd.Flags().String("max-results", "", "The maximum number of transaction events to list.")
		managedblockchainQuery_listFilteredTransactionEventsCmd.Flags().String("network", "", "The blockchain network where the transaction occurred.")
		managedblockchainQuery_listFilteredTransactionEventsCmd.Flags().String("next-token", "", "The pagination token that indicates the next set of results to retrieve.")
		managedblockchainQuery_listFilteredTransactionEventsCmd.Flags().String("sort", "", "The order by which the results will be sorted.")
		managedblockchainQuery_listFilteredTransactionEventsCmd.Flags().String("time-filter", "", "This container specifies the time frame for the transaction events returned in the response.")
		managedblockchainQuery_listFilteredTransactionEventsCmd.Flags().String("vout-filter", "", "This container specifies filtering attributes related to BITCOIN\\_VOUT event types")
		managedblockchainQuery_listFilteredTransactionEventsCmd.MarkFlagRequired("address-identifier-filter")
		managedblockchainQuery_listFilteredTransactionEventsCmd.MarkFlagRequired("network")
	})
	managedblockchainQueryCmd.AddCommand(managedblockchainQuery_listFilteredTransactionEventsCmd)
}
