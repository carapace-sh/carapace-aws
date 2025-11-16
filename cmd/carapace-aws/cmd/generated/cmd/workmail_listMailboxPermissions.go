package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_listMailboxPermissionsCmd = &cobra.Command{
	Use:   "list-mailbox-permissions",
	Short: "Lists the mailbox permissions associated with a user, group, or resource mailbox.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_listMailboxPermissionsCmd).Standalone()

	workmail_listMailboxPermissionsCmd.Flags().String("entity-id", "", "The identifier of the user, or resource for which to list mailbox permissions.")
	workmail_listMailboxPermissionsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	workmail_listMailboxPermissionsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	workmail_listMailboxPermissionsCmd.Flags().String("organization-id", "", "The identifier of the organization under which the user, group, or resource exists.")
	workmail_listMailboxPermissionsCmd.MarkFlagRequired("entity-id")
	workmail_listMailboxPermissionsCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_listMailboxPermissionsCmd)
}
