package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_listGroupMembersCmd = &cobra.Command{
	Use:   "list-group-members",
	Short: "Returns an overview of the members of a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_listGroupMembersCmd).Standalone()

	workmail_listGroupMembersCmd.Flags().String("group-id", "", "The identifier for the group to which the members (users or groups) are associated.")
	workmail_listGroupMembersCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	workmail_listGroupMembersCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	workmail_listGroupMembersCmd.Flags().String("organization-id", "", "The identifier for the organization under which the group exists.")
	workmail_listGroupMembersCmd.MarkFlagRequired("group-id")
	workmail_listGroupMembersCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_listGroupMembersCmd)
}
