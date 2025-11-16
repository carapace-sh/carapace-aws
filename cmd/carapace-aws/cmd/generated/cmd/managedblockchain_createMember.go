package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_createMemberCmd = &cobra.Command{
	Use:   "create-member",
	Short: "Creates a member within a Managed Blockchain network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_createMemberCmd).Standalone()

	managedblockchain_createMemberCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the operation.")
	managedblockchain_createMemberCmd.Flags().String("invitation-id", "", "The unique identifier of the invitation that is sent to the member to join the network.")
	managedblockchain_createMemberCmd.Flags().String("member-configuration", "", "Member configuration parameters.")
	managedblockchain_createMemberCmd.Flags().String("network-id", "", "The unique identifier of the network in which the member is created.")
	managedblockchain_createMemberCmd.MarkFlagRequired("client-request-token")
	managedblockchain_createMemberCmd.MarkFlagRequired("invitation-id")
	managedblockchain_createMemberCmd.MarkFlagRequired("member-configuration")
	managedblockchain_createMemberCmd.MarkFlagRequired("network-id")
	managedblockchainCmd.AddCommand(managedblockchain_createMemberCmd)
}
