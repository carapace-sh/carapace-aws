package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_updateCoreNetworkCmd = &cobra.Command{
	Use:   "update-core-network",
	Short: "Updates the description of a core network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_updateCoreNetworkCmd).Standalone()

	networkmanager_updateCoreNetworkCmd.Flags().String("core-network-id", "", "The ID of a core network.")
	networkmanager_updateCoreNetworkCmd.Flags().String("description", "", "The description of the update.")
	networkmanager_updateCoreNetworkCmd.MarkFlagRequired("core-network-id")
	networkmanagerCmd.AddCommand(networkmanager_updateCoreNetworkCmd)
}
