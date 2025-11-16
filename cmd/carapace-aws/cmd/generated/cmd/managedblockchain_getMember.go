package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_getMemberCmd = &cobra.Command{
	Use:   "get-member",
	Short: "Returns detailed information about a member.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_getMemberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(managedblockchain_getMemberCmd).Standalone()

		managedblockchain_getMemberCmd.Flags().String("member-id", "", "The unique identifier of the member.")
		managedblockchain_getMemberCmd.Flags().String("network-id", "", "The unique identifier of the network to which the member belongs.")
		managedblockchain_getMemberCmd.MarkFlagRequired("member-id")
		managedblockchain_getMemberCmd.MarkFlagRequired("network-id")
	})
	managedblockchainCmd.AddCommand(managedblockchain_getMemberCmd)
}
