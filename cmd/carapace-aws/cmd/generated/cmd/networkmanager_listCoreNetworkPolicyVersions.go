package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_listCoreNetworkPolicyVersionsCmd = &cobra.Command{
	Use:   "list-core-network-policy-versions",
	Short: "Returns a list of core network policy versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_listCoreNetworkPolicyVersionsCmd).Standalone()

	networkmanager_listCoreNetworkPolicyVersionsCmd.Flags().String("core-network-id", "", "The ID of a core network.")
	networkmanager_listCoreNetworkPolicyVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	networkmanager_listCoreNetworkPolicyVersionsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	networkmanager_listCoreNetworkPolicyVersionsCmd.MarkFlagRequired("core-network-id")
	networkmanagerCmd.AddCommand(networkmanager_listCoreNetworkPolicyVersionsCmd)
}
