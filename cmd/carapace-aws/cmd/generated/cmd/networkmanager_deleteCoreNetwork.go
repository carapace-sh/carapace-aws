package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_deleteCoreNetworkCmd = &cobra.Command{
	Use:   "delete-core-network",
	Short: "Deletes a core network along with all core network policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_deleteCoreNetworkCmd).Standalone()

	networkmanager_deleteCoreNetworkCmd.Flags().String("core-network-id", "", "The network ID of the deleted core network.")
	networkmanager_deleteCoreNetworkCmd.MarkFlagRequired("core-network-id")
	networkmanagerCmd.AddCommand(networkmanager_deleteCoreNetworkCmd)
}
