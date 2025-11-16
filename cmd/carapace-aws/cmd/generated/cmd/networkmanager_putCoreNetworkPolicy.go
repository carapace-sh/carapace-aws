package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_putCoreNetworkPolicyCmd = &cobra.Command{
	Use:   "put-core-network-policy",
	Short: "Creates a new, immutable version of a core network policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_putCoreNetworkPolicyCmd).Standalone()

	networkmanager_putCoreNetworkPolicyCmd.Flags().String("client-token", "", "The client token associated with the request.")
	networkmanager_putCoreNetworkPolicyCmd.Flags().String("core-network-id", "", "The ID of a core network.")
	networkmanager_putCoreNetworkPolicyCmd.Flags().String("description", "", "a core network policy description.")
	networkmanager_putCoreNetworkPolicyCmd.Flags().String("latest-version-id", "", "The ID of a core network policy.")
	networkmanager_putCoreNetworkPolicyCmd.Flags().String("policy-document", "", "The policy document.")
	networkmanager_putCoreNetworkPolicyCmd.MarkFlagRequired("core-network-id")
	networkmanager_putCoreNetworkPolicyCmd.MarkFlagRequired("policy-document")
	networkmanagerCmd.AddCommand(networkmanager_putCoreNetworkPolicyCmd)
}
