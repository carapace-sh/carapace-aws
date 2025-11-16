package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getCoreNetworkCmd = &cobra.Command{
	Use:   "get-core-network",
	Short: "Returns information about the LIVE policy for a core network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getCoreNetworkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_getCoreNetworkCmd).Standalone()

		networkmanager_getCoreNetworkCmd.Flags().String("core-network-id", "", "The ID of a core network.")
		networkmanager_getCoreNetworkCmd.MarkFlagRequired("core-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_getCoreNetworkCmd)
}
