package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_deleteGlobalNetworkCmd = &cobra.Command{
	Use:   "delete-global-network",
	Short: "Deletes an existing global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_deleteGlobalNetworkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_deleteGlobalNetworkCmd).Standalone()

		networkmanager_deleteGlobalNetworkCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_deleteGlobalNetworkCmd.MarkFlagRequired("global-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_deleteGlobalNetworkCmd)
}
