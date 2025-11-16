package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_createProposalCmd = &cobra.Command{
	Use:   "create-proposal",
	Short: "Creates a proposal for a change to the network that other members of the network can vote on, for example, a proposal to add a new member to the network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_createProposalCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(managedblockchain_createProposalCmd).Standalone()

		managedblockchain_createProposalCmd.Flags().String("actions", "", "The type of actions proposed, such as inviting a member or removing a member.")
		managedblockchain_createProposalCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the operation.")
		managedblockchain_createProposalCmd.Flags().String("description", "", "A description for the proposal that is visible to voting members, for example, \"Proposal to add Example Corp. as member.\"")
		managedblockchain_createProposalCmd.Flags().String("member-id", "", "The unique identifier of the member that is creating the proposal.")
		managedblockchain_createProposalCmd.Flags().String("network-id", "", "The unique identifier of the network for which the proposal is made.")
		managedblockchain_createProposalCmd.Flags().String("tags", "", "Tags to assign to the proposal.")
		managedblockchain_createProposalCmd.MarkFlagRequired("actions")
		managedblockchain_createProposalCmd.MarkFlagRequired("client-request-token")
		managedblockchain_createProposalCmd.MarkFlagRequired("member-id")
		managedblockchain_createProposalCmd.MarkFlagRequired("network-id")
	})
	managedblockchainCmd.AddCommand(managedblockchain_createProposalCmd)
}
