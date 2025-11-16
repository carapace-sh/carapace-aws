package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_associateMemberToGroupCmd = &cobra.Command{
	Use:   "associate-member-to-group",
	Short: "Adds a member (user or group) to the group's set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_associateMemberToGroupCmd).Standalone()

	workmail_associateMemberToGroupCmd.Flags().String("group-id", "", "The group to which the member (user or group) is associated.")
	workmail_associateMemberToGroupCmd.Flags().String("member-id", "", "The member (user or group) to associate to the group.")
	workmail_associateMemberToGroupCmd.Flags().String("organization-id", "", "The organization under which the group exists.")
	workmail_associateMemberToGroupCmd.MarkFlagRequired("group-id")
	workmail_associateMemberToGroupCmd.MarkFlagRequired("member-id")
	workmail_associateMemberToGroupCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_associateMemberToGroupCmd)
}
