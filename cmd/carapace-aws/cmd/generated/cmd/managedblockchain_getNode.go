package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_getNodeCmd = &cobra.Command{
	Use:   "get-node",
	Short: "Returns detailed information about a node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_getNodeCmd).Standalone()

	managedblockchain_getNodeCmd.Flags().String("member-id", "", "The unique identifier of the member that owns the node.")
	managedblockchain_getNodeCmd.Flags().String("network-id", "", "The unique identifier of the network that the node is on.")
	managedblockchain_getNodeCmd.Flags().String("node-id", "", "The unique identifier of the node.")
	managedblockchain_getNodeCmd.MarkFlagRequired("network-id")
	managedblockchain_getNodeCmd.MarkFlagRequired("node-id")
	managedblockchainCmd.AddCommand(managedblockchain_getNodeCmd)
}
