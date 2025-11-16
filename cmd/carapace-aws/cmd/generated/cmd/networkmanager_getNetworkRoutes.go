package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getNetworkRoutesCmd = &cobra.Command{
	Use:   "get-network-routes",
	Short: "Gets the network routes of the specified global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getNetworkRoutesCmd).Standalone()

	networkmanager_getNetworkRoutesCmd.Flags().String("destination-filters", "", "Filter by route table destination.")
	networkmanager_getNetworkRoutesCmd.Flags().String("exact-cidr-matches", "", "An exact CIDR block.")
	networkmanager_getNetworkRoutesCmd.Flags().String("global-network-id", "", "The ID of the global network.")
	networkmanager_getNetworkRoutesCmd.Flags().String("longest-prefix-matches", "", "The most specific route that matches the traffic (longest prefix match).")
	networkmanager_getNetworkRoutesCmd.Flags().String("prefix-list-ids", "", "The IDs of the prefix lists.")
	networkmanager_getNetworkRoutesCmd.Flags().String("route-table-identifier", "", "The ID of the route table.")
	networkmanager_getNetworkRoutesCmd.Flags().String("states", "", "The route states.")
	networkmanager_getNetworkRoutesCmd.Flags().String("subnet-of-matches", "", "The routes with a subnet that match the specified CIDR filter.")
	networkmanager_getNetworkRoutesCmd.Flags().String("supernet-of-matches", "", "The routes with a CIDR that encompasses the CIDR filter.")
	networkmanager_getNetworkRoutesCmd.Flags().String("types", "", "The route types.")
	networkmanager_getNetworkRoutesCmd.MarkFlagRequired("global-network-id")
	networkmanager_getNetworkRoutesCmd.MarkFlagRequired("route-table-identifier")
	networkmanagerCmd.AddCommand(networkmanager_getNetworkRoutesCmd)
}
