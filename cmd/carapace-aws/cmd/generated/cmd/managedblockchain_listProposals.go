package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_listProposalsCmd = &cobra.Command{
	Use:   "list-proposals",
	Short: "Returns a list of proposals for the network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_listProposalsCmd).Standalone()

	managedblockchain_listProposalsCmd.Flags().String("max-results", "", "The maximum number of proposals to return.")
	managedblockchain_listProposalsCmd.Flags().String("network-id", "", "The unique identifier of the network.")
	managedblockchain_listProposalsCmd.Flags().String("next-token", "", "The pagination token that indicates the next set of results to retrieve.")
	managedblockchain_listProposalsCmd.MarkFlagRequired("network-id")
	managedblockchainCmd.AddCommand(managedblockchain_listProposalsCmd)
}
