package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_createNetworkCmd = &cobra.Command{
	Use:   "create-network",
	Short: "Creates a new blockchain network using Amazon Managed Blockchain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_createNetworkCmd).Standalone()

	managedblockchain_createNetworkCmd.Flags().String("client-request-token", "", "This is a unique, case-sensitive identifier that you provide to ensure the idempotency of the operation.")
	managedblockchain_createNetworkCmd.Flags().String("description", "", "An optional description for the network.")
	managedblockchain_createNetworkCmd.Flags().String("framework", "", "The blockchain framework that the network uses.")
	managedblockchain_createNetworkCmd.Flags().String("framework-configuration", "", "Configuration properties of the blockchain framework relevant to the network configuration.")
	managedblockchain_createNetworkCmd.Flags().String("framework-version", "", "The version of the blockchain framework that the network uses.")
	managedblockchain_createNetworkCmd.Flags().String("member-configuration", "", "Configuration properties for the first member within the network.")
	managedblockchain_createNetworkCmd.Flags().String("name", "", "The name of the network.")
	managedblockchain_createNetworkCmd.Flags().String("tags", "", "Tags to assign to the network.")
	managedblockchain_createNetworkCmd.Flags().String("voting-policy", "", "The voting rules used by the network to determine if a proposal is approved.")
	managedblockchain_createNetworkCmd.MarkFlagRequired("client-request-token")
	managedblockchain_createNetworkCmd.MarkFlagRequired("framework")
	managedblockchain_createNetworkCmd.MarkFlagRequired("framework-version")
	managedblockchain_createNetworkCmd.MarkFlagRequired("member-configuration")
	managedblockchain_createNetworkCmd.MarkFlagRequired("name")
	managedblockchain_createNetworkCmd.MarkFlagRequired("voting-policy")
	managedblockchainCmd.AddCommand(managedblockchain_createNetworkCmd)
}
