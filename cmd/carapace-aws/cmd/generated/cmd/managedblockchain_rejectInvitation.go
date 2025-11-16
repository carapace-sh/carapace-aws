package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_rejectInvitationCmd = &cobra.Command{
	Use:   "reject-invitation",
	Short: "Rejects an invitation to join a network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_rejectInvitationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(managedblockchain_rejectInvitationCmd).Standalone()

		managedblockchain_rejectInvitationCmd.Flags().String("invitation-id", "", "The unique identifier of the invitation to reject.")
		managedblockchain_rejectInvitationCmd.MarkFlagRequired("invitation-id")
	})
	managedblockchainCmd.AddCommand(managedblockchain_rejectInvitationCmd)
}
