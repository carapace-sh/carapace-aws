package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_voteOnProposalCmd = &cobra.Command{
	Use:   "vote-on-proposal",
	Short: "Casts a vote for a specified `ProposalId` on behalf of a member.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_voteOnProposalCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(managedblockchain_voteOnProposalCmd).Standalone()

		managedblockchain_voteOnProposalCmd.Flags().String("network-id", "", "The unique identifier of the network.")
		managedblockchain_voteOnProposalCmd.Flags().String("proposal-id", "", "The unique identifier of the proposal.")
		managedblockchain_voteOnProposalCmd.Flags().String("vote", "", "The value of the vote.")
		managedblockchain_voteOnProposalCmd.Flags().String("voter-member-id", "", "The unique identifier of the member casting the vote.")
		managedblockchain_voteOnProposalCmd.MarkFlagRequired("network-id")
		managedblockchain_voteOnProposalCmd.MarkFlagRequired("proposal-id")
		managedblockchain_voteOnProposalCmd.MarkFlagRequired("vote")
		managedblockchain_voteOnProposalCmd.MarkFlagRequired("voter-member-id")
	})
	managedblockchainCmd.AddCommand(managedblockchain_voteOnProposalCmd)
}
