package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getCoreNetworkPolicyCmd = &cobra.Command{
	Use:   "get-core-network-policy",
	Short: "Returns details about a core network policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getCoreNetworkPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_getCoreNetworkPolicyCmd).Standalone()

		networkmanager_getCoreNetworkPolicyCmd.Flags().String("alias", "", "The alias of a core network policy")
		networkmanager_getCoreNetworkPolicyCmd.Flags().String("core-network-id", "", "The ID of a core network.")
		networkmanager_getCoreNetworkPolicyCmd.Flags().String("policy-version-id", "", "The ID of a core network policy version.")
		networkmanager_getCoreNetworkPolicyCmd.MarkFlagRequired("core-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_getCoreNetworkPolicyCmd)
}
