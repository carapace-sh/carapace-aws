package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_listCoreNetworksCmd = &cobra.Command{
	Use:   "list-core-networks",
	Short: "Returns a list of owned and shared core networks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_listCoreNetworksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_listCoreNetworksCmd).Standalone()

		networkmanager_listCoreNetworksCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		networkmanager_listCoreNetworksCmd.Flags().String("next-token", "", "The token for the next page of results.")
	})
	networkmanagerCmd.AddCommand(networkmanager_listCoreNetworksCmd)
}
