package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_deleteMemberCmd = &cobra.Command{
	Use:   "delete-member",
	Short: "Deletes a member.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_deleteMemberCmd).Standalone()

	managedblockchain_deleteMemberCmd.Flags().String("member-id", "", "The unique identifier of the member to remove.")
	managedblockchain_deleteMemberCmd.Flags().String("network-id", "", "The unique identifier of the network from which the member is removed.")
	managedblockchain_deleteMemberCmd.MarkFlagRequired("member-id")
	managedblockchain_deleteMemberCmd.MarkFlagRequired("network-id")
	managedblockchainCmd.AddCommand(managedblockchain_deleteMemberCmd)
}
