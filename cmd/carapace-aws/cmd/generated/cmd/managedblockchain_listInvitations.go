package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_listInvitationsCmd = &cobra.Command{
	Use:   "list-invitations",
	Short: "Returns a list of all invitations for the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_listInvitationsCmd).Standalone()

	managedblockchain_listInvitationsCmd.Flags().String("max-results", "", "The maximum number of invitations to return.")
	managedblockchain_listInvitationsCmd.Flags().String("next-token", "", "The pagination token that indicates the next set of results to retrieve.")
	managedblockchainCmd.AddCommand(managedblockchain_listInvitationsCmd)
}
