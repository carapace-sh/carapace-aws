package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_createAccessorCmd = &cobra.Command{
	Use:   "create-accessor",
	Short: "Creates a new accessor for use with Amazon Managed Blockchain service that supports token based access.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_createAccessorCmd).Standalone()

	managedblockchain_createAccessorCmd.Flags().String("accessor-type", "", "The type of accessor.")
	managedblockchain_createAccessorCmd.Flags().String("client-request-token", "", "This is a unique, case-sensitive identifier that you provide to ensure the idempotency of the operation.")
	managedblockchain_createAccessorCmd.Flags().String("network-type", "", "The blockchain network that the `Accessor` token is created for.")
	managedblockchain_createAccessorCmd.Flags().String("tags", "", "Tags to assign to the Accessor.")
	managedblockchain_createAccessorCmd.MarkFlagRequired("accessor-type")
	managedblockchain_createAccessorCmd.MarkFlagRequired("client-request-token")
	managedblockchainCmd.AddCommand(managedblockchain_createAccessorCmd)
}
