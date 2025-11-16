package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getNetworkResourceCountsCmd = &cobra.Command{
	Use:   "get-network-resource-counts",
	Short: "Gets the count of network resources, by resource type, for the specified global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getNetworkResourceCountsCmd).Standalone()

	networkmanager_getNetworkResourceCountsCmd.Flags().String("global-network-id", "", "The ID of the global network.")
	networkmanager_getNetworkResourceCountsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	networkmanager_getNetworkResourceCountsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	networkmanager_getNetworkResourceCountsCmd.Flags().String("resource-type", "", "The resource type.")
	networkmanager_getNetworkResourceCountsCmd.MarkFlagRequired("global-network-id")
	networkmanagerCmd.AddCommand(networkmanager_getNetworkResourceCountsCmd)
}
