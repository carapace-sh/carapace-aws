package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_getMailboxDetailsCmd = &cobra.Command{
	Use:   "get-mailbox-details",
	Short: "Requests a user's mailbox details for a specified organization and user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_getMailboxDetailsCmd).Standalone()

	workmail_getMailboxDetailsCmd.Flags().String("organization-id", "", "The identifier for the organization that contains the user whose mailbox details are being requested.")
	workmail_getMailboxDetailsCmd.Flags().String("user-id", "", "The identifier for the user whose mailbox details are being requested.")
	workmail_getMailboxDetailsCmd.MarkFlagRequired("organization-id")
	workmail_getMailboxDetailsCmd.MarkFlagRequired("user-id")
	workmailCmd.AddCommand(workmail_getMailboxDetailsCmd)
}
