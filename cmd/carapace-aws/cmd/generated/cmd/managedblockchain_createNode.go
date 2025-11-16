package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_createNodeCmd = &cobra.Command{
	Use:   "create-node",
	Short: "Creates a node on the specified blockchain network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_createNodeCmd).Standalone()

	managedblockchain_createNodeCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the operation.")
	managedblockchain_createNodeCmd.Flags().String("member-id", "", "The unique identifier of the member that owns this node.")
	managedblockchain_createNodeCmd.Flags().String("network-id", "", "The unique identifier of the network for the node.")
	managedblockchain_createNodeCmd.Flags().String("node-configuration", "", "The properties of a node configuration.")
	managedblockchain_createNodeCmd.Flags().String("tags", "", "Tags to assign to the node.")
	managedblockchain_createNodeCmd.MarkFlagRequired("client-request-token")
	managedblockchain_createNodeCmd.MarkFlagRequired("network-id")
	managedblockchain_createNodeCmd.MarkFlagRequired("node-configuration")
	managedblockchainCmd.AddCommand(managedblockchain_createNodeCmd)
}
