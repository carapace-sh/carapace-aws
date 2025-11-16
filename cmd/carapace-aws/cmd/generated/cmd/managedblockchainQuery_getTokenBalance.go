package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchainQuery_getTokenBalanceCmd = &cobra.Command{
	Use:   "get-token-balance",
	Short: "Gets the balance of a specific token, including native tokens, for a given address (wallet or contract) on the blockchain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchainQuery_getTokenBalanceCmd).Standalone()

	managedblockchainQuery_getTokenBalanceCmd.Flags().String("at-blockchain-instant", "", "The time for when the TokenBalance is requested or the current time if a time is not provided in the request.")
	managedblockchainQuery_getTokenBalanceCmd.Flags().String("owner-identifier", "", "The container for the identifier for the owner.")
	managedblockchainQuery_getTokenBalanceCmd.Flags().String("token-identifier", "", "The container for the identifier for the token, including the unique token ID and its blockchain network.")
	managedblockchainQuery_getTokenBalanceCmd.MarkFlagRequired("owner-identifier")
	managedblockchainQuery_getTokenBalanceCmd.MarkFlagRequired("token-identifier")
	managedblockchainQueryCmd.AddCommand(managedblockchainQuery_getTokenBalanceCmd)
}
