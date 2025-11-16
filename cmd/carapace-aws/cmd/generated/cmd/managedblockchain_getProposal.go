package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_getProposalCmd = &cobra.Command{
	Use:   "get-proposal",
	Short: "Returns detailed information about a proposal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_getProposalCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(managedblockchain_getProposalCmd).Standalone()

		managedblockchain_getProposalCmd.Flags().String("network-id", "", "The unique identifier of the network for which the proposal is made.")
		managedblockchain_getProposalCmd.Flags().String("proposal-id", "", "The unique identifier of the proposal.")
		managedblockchain_getProposalCmd.MarkFlagRequired("network-id")
		managedblockchain_getProposalCmd.MarkFlagRequired("proposal-id")
	})
	managedblockchainCmd.AddCommand(managedblockchain_getProposalCmd)
}
