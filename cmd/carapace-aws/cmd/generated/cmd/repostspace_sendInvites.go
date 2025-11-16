package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_sendInvitesCmd = &cobra.Command{
	Use:   "send-invites",
	Short: "Sends an invitation email to selected users and groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_sendInvitesCmd).Standalone()

	repostspace_sendInvitesCmd.Flags().String("accessor-ids", "", "The array of identifiers for the users and groups.")
	repostspace_sendInvitesCmd.Flags().String("body", "", "The body of the invite.")
	repostspace_sendInvitesCmd.Flags().String("space-id", "", "The ID of the private re:Post.")
	repostspace_sendInvitesCmd.Flags().String("title", "", "The title of the invite.")
	repostspace_sendInvitesCmd.MarkFlagRequired("accessor-ids")
	repostspace_sendInvitesCmd.MarkFlagRequired("body")
	repostspace_sendInvitesCmd.MarkFlagRequired("space-id")
	repostspace_sendInvitesCmd.MarkFlagRequired("title")
	repostspaceCmd.AddCommand(repostspace_sendInvitesCmd)
}
