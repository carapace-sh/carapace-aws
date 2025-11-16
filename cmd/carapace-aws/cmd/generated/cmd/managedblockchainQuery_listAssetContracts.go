package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchainQuery_listAssetContractsCmd = &cobra.Command{
	Use:   "list-asset-contracts",
	Short: "Lists all the contracts for a given contract type deployed by an address (either a contract address or a wallet address).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchainQuery_listAssetContractsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(managedblockchainQuery_listAssetContractsCmd).Standalone()

		managedblockchainQuery_listAssetContractsCmd.Flags().String("contract-filter", "", "Contains the filter parameter for the request.")
		managedblockchainQuery_listAssetContractsCmd.Flags().String("max-results", "", "The maximum number of contracts to list.")
		managedblockchainQuery_listAssetContractsCmd.Flags().String("next-token", "", "The pagination token that indicates the next set of results to retrieve.")
		managedblockchainQuery_listAssetContractsCmd.MarkFlagRequired("contract-filter")
	})
	managedblockchainQueryCmd.AddCommand(managedblockchainQuery_listAssetContractsCmd)
}
