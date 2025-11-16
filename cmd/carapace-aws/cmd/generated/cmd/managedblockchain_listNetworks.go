package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_listNetworksCmd = &cobra.Command{
	Use:   "list-networks",
	Short: "Returns information about the networks in which the current Amazon Web Services account participates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_listNetworksCmd).Standalone()

	managedblockchain_listNetworksCmd.Flags().String("framework", "", "An optional framework specifier.")
	managedblockchain_listNetworksCmd.Flags().String("max-results", "", "The maximum number of networks to list.")
	managedblockchain_listNetworksCmd.Flags().String("name", "", "The name of the network.")
	managedblockchain_listNetworksCmd.Flags().String("next-token", "", "The pagination token that indicates the next set of results to retrieve.")
	managedblockchain_listNetworksCmd.Flags().String("status", "", "An optional status specifier.")
	managedblockchainCmd.AddCommand(managedblockchain_listNetworksCmd)
}
