package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchainQuery_listTokenBalancesCmd = &cobra.Command{
	Use:   "list-token-balances",
	Short: "This action returns the following for a given blockchain network:\n\n- Lists all token balances owned by an address (either a contract address or a wallet address).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchainQuery_listTokenBalancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(managedblockchainQuery_listTokenBalancesCmd).Standalone()

		managedblockchainQuery_listTokenBalancesCmd.Flags().String("max-results", "", "The maximum number of token balances to return.")
		managedblockchainQuery_listTokenBalancesCmd.Flags().String("next-token", "", "The pagination token that indicates the next set of results to retrieve.")
		managedblockchainQuery_listTokenBalancesCmd.Flags().String("owner-filter", "", "The contract or wallet address on the blockchain network by which to filter the request.")
		managedblockchainQuery_listTokenBalancesCmd.Flags().String("token-filter", "", "The contract address or a token identifier on the blockchain network by which to filter the request.")
		managedblockchainQuery_listTokenBalancesCmd.MarkFlagRequired("token-filter")
	})
	managedblockchainQueryCmd.AddCommand(managedblockchainQuery_listTokenBalancesCmd)
}
