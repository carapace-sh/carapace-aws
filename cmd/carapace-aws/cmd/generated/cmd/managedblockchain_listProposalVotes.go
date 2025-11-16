package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_listProposalVotesCmd = &cobra.Command{
	Use:   "list-proposal-votes",
	Short: "Returns the list of votes for a specified proposal, including the value of each vote and the unique identifier of the member that cast the vote.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_listProposalVotesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(managedblockchain_listProposalVotesCmd).Standalone()

		managedblockchain_listProposalVotesCmd.Flags().String("max-results", "", "The maximum number of votes to return.")
		managedblockchain_listProposalVotesCmd.Flags().String("network-id", "", "The unique identifier of the network.")
		managedblockchain_listProposalVotesCmd.Flags().String("next-token", "", "The pagination token that indicates the next set of results to retrieve.")
		managedblockchain_listProposalVotesCmd.Flags().String("proposal-id", "", "The unique identifier of the proposal.")
		managedblockchain_listProposalVotesCmd.MarkFlagRequired("network-id")
		managedblockchain_listProposalVotesCmd.MarkFlagRequired("proposal-id")
	})
	managedblockchainCmd.AddCommand(managedblockchain_listProposalVotesCmd)
}
