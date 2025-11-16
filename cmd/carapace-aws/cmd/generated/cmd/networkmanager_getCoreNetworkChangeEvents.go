package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getCoreNetworkChangeEventsCmd = &cobra.Command{
	Use:   "get-core-network-change-events",
	Short: "Returns information about a core network change event.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getCoreNetworkChangeEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_getCoreNetworkChangeEventsCmd).Standalone()

		networkmanager_getCoreNetworkChangeEventsCmd.Flags().String("core-network-id", "", "The ID of a core network.")
		networkmanager_getCoreNetworkChangeEventsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		networkmanager_getCoreNetworkChangeEventsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		networkmanager_getCoreNetworkChangeEventsCmd.Flags().String("policy-version-id", "", "The ID of the policy version.")
		networkmanager_getCoreNetworkChangeEventsCmd.MarkFlagRequired("core-network-id")
		networkmanager_getCoreNetworkChangeEventsCmd.MarkFlagRequired("policy-version-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_getCoreNetworkChangeEventsCmd)
}
