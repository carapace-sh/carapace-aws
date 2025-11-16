package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_executeCoreNetworkChangeSetCmd = &cobra.Command{
	Use:   "execute-core-network-change-set",
	Short: "Executes a change set on your core network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_executeCoreNetworkChangeSetCmd).Standalone()

	networkmanager_executeCoreNetworkChangeSetCmd.Flags().String("core-network-id", "", "The ID of a core network.")
	networkmanager_executeCoreNetworkChangeSetCmd.Flags().String("policy-version-id", "", "The ID of the policy version.")
	networkmanager_executeCoreNetworkChangeSetCmd.MarkFlagRequired("core-network-id")
	networkmanager_executeCoreNetworkChangeSetCmd.MarkFlagRequired("policy-version-id")
	networkmanagerCmd.AddCommand(networkmanager_executeCoreNetworkChangeSetCmd)
}
