package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchainQuery_batchGetTokenBalanceCmd = &cobra.Command{
	Use:   "batch-get-token-balance",
	Short: "Gets the token balance for a batch of tokens by using the `BatchGetTokenBalance` action for every token in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchainQuery_batchGetTokenBalanceCmd).Standalone()

	managedblockchainQuery_batchGetTokenBalanceCmd.Flags().String("get-token-balance-inputs", "", "An array of `BatchGetTokenBalanceInputItem` objects whose balance is being requested.")
	managedblockchainQueryCmd.AddCommand(managedblockchainQuery_batchGetTokenBalanceCmd)
}
