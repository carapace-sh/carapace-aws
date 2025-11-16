package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_updateNodeCmd = &cobra.Command{
	Use:   "update-node",
	Short: "Updates a node configuration with new parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_updateNodeCmd).Standalone()

	managedblockchain_updateNodeCmd.Flags().String("log-publishing-configuration", "", "Configuration properties for publishing to Amazon CloudWatch Logs.")
	managedblockchain_updateNodeCmd.Flags().String("member-id", "", "The unique identifier of the member that owns the node.")
	managedblockchain_updateNodeCmd.Flags().String("network-id", "", "The unique identifier of the network that the node is on.")
	managedblockchain_updateNodeCmd.Flags().String("node-id", "", "The unique identifier of the node.")
	managedblockchain_updateNodeCmd.MarkFlagRequired("network-id")
	managedblockchain_updateNodeCmd.MarkFlagRequired("node-id")
	managedblockchainCmd.AddCommand(managedblockchain_updateNodeCmd)
}
