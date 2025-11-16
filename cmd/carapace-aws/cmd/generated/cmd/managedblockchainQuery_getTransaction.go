package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchainQuery_getTransactionCmd = &cobra.Command{
	Use:   "get-transaction",
	Short: "Gets the details of a transaction.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchainQuery_getTransactionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(managedblockchainQuery_getTransactionCmd).Standalone()

		managedblockchainQuery_getTransactionCmd.Flags().String("network", "", "The blockchain network where the transaction occurred.")
		managedblockchainQuery_getTransactionCmd.Flags().String("transaction-hash", "", "The hash of a transaction.")
		managedblockchainQuery_getTransactionCmd.Flags().String("transaction-id", "", "The identifier of a Bitcoin transaction.")
		managedblockchainQuery_getTransactionCmd.MarkFlagRequired("network")
	})
	managedblockchainQueryCmd.AddCommand(managedblockchainQuery_getTransactionCmd)
}
