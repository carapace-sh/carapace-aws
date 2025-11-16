package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_createCoreNetworkCmd = &cobra.Command{
	Use:   "create-core-network",
	Short: "Creates a core network as part of your global network, and optionally, with a core network policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_createCoreNetworkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_createCoreNetworkCmd).Standalone()

		networkmanager_createCoreNetworkCmd.Flags().String("client-token", "", "The client token associated with a core network request.")
		networkmanager_createCoreNetworkCmd.Flags().String("description", "", "The description of a core network.")
		networkmanager_createCoreNetworkCmd.Flags().String("global-network-id", "", "The ID of the global network that a core network will be a part of.")
		networkmanager_createCoreNetworkCmd.Flags().String("policy-document", "", "The policy document for creating a core network.")
		networkmanager_createCoreNetworkCmd.Flags().String("tags", "", "Key-value tags associated with a core network request.")
		networkmanager_createCoreNetworkCmd.MarkFlagRequired("global-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_createCoreNetworkCmd)
}
