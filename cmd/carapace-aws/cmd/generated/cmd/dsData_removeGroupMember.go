package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsData_removeGroupMemberCmd = &cobra.Command{
	Use:   "remove-group-member",
	Short: "Removes a member from a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsData_removeGroupMemberCmd).Standalone()

	dsData_removeGroupMemberCmd.Flags().String("client-token", "", "A unique and case-sensitive identifier that you provide to make sure the idempotency of the request, so multiple identical calls have the same effect as one single call.")
	dsData_removeGroupMemberCmd.Flags().String("directory-id", "", "The identifier (ID) of the directory that's associated with the member.")
	dsData_removeGroupMemberCmd.Flags().String("group-name", "", "The name of the group.")
	dsData_removeGroupMemberCmd.Flags().String("member-name", "", "The `SAMAccountName` of the user, group, or computer to remove from the group.")
	dsData_removeGroupMemberCmd.Flags().String("member-realm", "", "The domain name that's associated with the group member.")
	dsData_removeGroupMemberCmd.MarkFlagRequired("directory-id")
	dsData_removeGroupMemberCmd.MarkFlagRequired("group-name")
	dsData_removeGroupMemberCmd.MarkFlagRequired("member-name")
	dsDataCmd.AddCommand(dsData_removeGroupMemberCmd)
}
