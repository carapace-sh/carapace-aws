package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_restoreCoreNetworkPolicyVersionCmd = &cobra.Command{
	Use:   "restore-core-network-policy-version",
	Short: "Restores a previous policy version as a new, immutable version of a core network policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_restoreCoreNetworkPolicyVersionCmd).Standalone()

	networkmanager_restoreCoreNetworkPolicyVersionCmd.Flags().String("core-network-id", "", "The ID of a core network.")
	networkmanager_restoreCoreNetworkPolicyVersionCmd.Flags().String("policy-version-id", "", "The ID of the policy version to restore.")
	networkmanager_restoreCoreNetworkPolicyVersionCmd.MarkFlagRequired("core-network-id")
	networkmanager_restoreCoreNetworkPolicyVersionCmd.MarkFlagRequired("policy-version-id")
	networkmanagerCmd.AddCommand(networkmanager_restoreCoreNetworkPolicyVersionCmd)
}
