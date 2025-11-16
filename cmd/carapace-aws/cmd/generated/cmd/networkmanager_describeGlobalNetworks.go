package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_describeGlobalNetworksCmd = &cobra.Command{
	Use:   "describe-global-networks",
	Short: "Describes one or more global networks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_describeGlobalNetworksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_describeGlobalNetworksCmd).Standalone()

		networkmanager_describeGlobalNetworksCmd.Flags().String("global-network-ids", "", "The IDs of one or more global networks.")
		networkmanager_describeGlobalNetworksCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		networkmanager_describeGlobalNetworksCmd.Flags().String("next-token", "", "The token for the next page of results.")
	})
	networkmanagerCmd.AddCommand(networkmanager_describeGlobalNetworksCmd)
}
