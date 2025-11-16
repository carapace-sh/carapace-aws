package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_listMembersCmd = &cobra.Command{
	Use:   "list-members",
	Short: "Returns a list of the members in a network and properties of their configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_listMembersCmd).Standalone()

	managedblockchain_listMembersCmd.Flags().String("is-owned", "", "An optional Boolean value.")
	managedblockchain_listMembersCmd.Flags().String("max-results", "", "The maximum number of members to return in the request.")
	managedblockchain_listMembersCmd.Flags().String("name", "", "The optional name of the member to list.")
	managedblockchain_listMembersCmd.Flags().String("network-id", "", "The unique identifier of the network for which to list members.")
	managedblockchain_listMembersCmd.Flags().String("next-token", "", "The pagination token that indicates the next set of results to retrieve.")
	managedblockchain_listMembersCmd.Flags().String("status", "", "An optional status specifier.")
	managedblockchain_listMembersCmd.MarkFlagRequired("network-id")
	managedblockchainCmd.AddCommand(managedblockchain_listMembersCmd)
}
