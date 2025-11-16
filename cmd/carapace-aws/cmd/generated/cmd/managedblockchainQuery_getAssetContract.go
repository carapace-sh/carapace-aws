package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchainQuery_getAssetContractCmd = &cobra.Command{
	Use:   "get-asset-contract",
	Short: "Gets the information about a specific contract deployed on the blockchain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchainQuery_getAssetContractCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(managedblockchainQuery_getAssetContractCmd).Standalone()

		managedblockchainQuery_getAssetContractCmd.Flags().String("contract-identifier", "", "Contains the blockchain address and network information about the contract.")
		managedblockchainQuery_getAssetContractCmd.MarkFlagRequired("contract-identifier")
	})
	managedblockchainQueryCmd.AddCommand(managedblockchainQuery_getAssetContractCmd)
}
