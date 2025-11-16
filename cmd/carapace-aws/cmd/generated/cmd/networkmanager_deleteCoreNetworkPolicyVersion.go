package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_deleteCoreNetworkPolicyVersionCmd = &cobra.Command{
	Use:   "delete-core-network-policy-version",
	Short: "Deletes a policy version from a core network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_deleteCoreNetworkPolicyVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_deleteCoreNetworkPolicyVersionCmd).Standalone()

		networkmanager_deleteCoreNetworkPolicyVersionCmd.Flags().String("core-network-id", "", "The ID of a core network for the deleted policy.")
		networkmanager_deleteCoreNetworkPolicyVersionCmd.Flags().String("policy-version-id", "", "The version ID of the deleted policy.")
		networkmanager_deleteCoreNetworkPolicyVersionCmd.MarkFlagRequired("core-network-id")
		networkmanager_deleteCoreNetworkPolicyVersionCmd.MarkFlagRequired("policy-version-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_deleteCoreNetworkPolicyVersionCmd)
}
