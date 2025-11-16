package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getCoreNetworkChangeSetCmd = &cobra.Command{
	Use:   "get-core-network-change-set",
	Short: "Returns a change set between the LIVE core network policy and a submitted policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getCoreNetworkChangeSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_getCoreNetworkChangeSetCmd).Standalone()

		networkmanager_getCoreNetworkChangeSetCmd.Flags().String("core-network-id", "", "The ID of a core network.")
		networkmanager_getCoreNetworkChangeSetCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		networkmanager_getCoreNetworkChangeSetCmd.Flags().String("next-token", "", "The token for the next page of results.")
		networkmanager_getCoreNetworkChangeSetCmd.Flags().String("policy-version-id", "", "The ID of the policy version.")
		networkmanager_getCoreNetworkChangeSetCmd.MarkFlagRequired("core-network-id")
		networkmanager_getCoreNetworkChangeSetCmd.MarkFlagRequired("policy-version-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_getCoreNetworkChangeSetCmd)
}
