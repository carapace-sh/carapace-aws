package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_updateGlobalNetworkCmd = &cobra.Command{
	Use:   "update-global-network",
	Short: "Updates an existing global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_updateGlobalNetworkCmd).Standalone()

	networkmanager_updateGlobalNetworkCmd.Flags().String("description", "", "A description of the global network.")
	networkmanager_updateGlobalNetworkCmd.Flags().String("global-network-id", "", "The ID of your global network.")
	networkmanager_updateGlobalNetworkCmd.MarkFlagRequired("global-network-id")
	networkmanagerCmd.AddCommand(networkmanager_updateGlobalNetworkCmd)
}
