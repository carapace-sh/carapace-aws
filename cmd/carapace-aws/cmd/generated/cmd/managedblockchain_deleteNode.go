package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_deleteNodeCmd = &cobra.Command{
	Use:   "delete-node",
	Short: "Deletes a node that your Amazon Web Services account owns.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_deleteNodeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(managedblockchain_deleteNodeCmd).Standalone()

		managedblockchain_deleteNodeCmd.Flags().String("member-id", "", "The unique identifier of the member that owns this node.")
		managedblockchain_deleteNodeCmd.Flags().String("network-id", "", "The unique identifier of the network that the node is on.")
		managedblockchain_deleteNodeCmd.Flags().String("node-id", "", "The unique identifier of the node.")
		managedblockchain_deleteNodeCmd.MarkFlagRequired("network-id")
		managedblockchain_deleteNodeCmd.MarkFlagRequired("node-id")
	})
	managedblockchainCmd.AddCommand(managedblockchain_deleteNodeCmd)
}
