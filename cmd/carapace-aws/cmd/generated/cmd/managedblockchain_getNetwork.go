package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_getNetworkCmd = &cobra.Command{
	Use:   "get-network",
	Short: "Returns detailed information about a network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_getNetworkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(managedblockchain_getNetworkCmd).Standalone()

		managedblockchain_getNetworkCmd.Flags().String("network-id", "", "The unique identifier of the network to get information about.")
		managedblockchain_getNetworkCmd.MarkFlagRequired("network-id")
	})
	managedblockchainCmd.AddCommand(managedblockchain_getNetworkCmd)
}
