package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_disassociateMemberFromGroupCmd = &cobra.Command{
	Use:   "disassociate-member-from-group",
	Short: "Removes a member from a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_disassociateMemberFromGroupCmd).Standalone()

	workmail_disassociateMemberFromGroupCmd.Flags().String("group-id", "", "The identifier for the group from which members are removed.")
	workmail_disassociateMemberFromGroupCmd.Flags().String("member-id", "", "The identifier for the member to be removed from the group.")
	workmail_disassociateMemberFromGroupCmd.Flags().String("organization-id", "", "The identifier for the organization under which the group exists.")
	workmail_disassociateMemberFromGroupCmd.MarkFlagRequired("group-id")
	workmail_disassociateMemberFromGroupCmd.MarkFlagRequired("member-id")
	workmail_disassociateMemberFromGroupCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_disassociateMemberFromGroupCmd)
}
