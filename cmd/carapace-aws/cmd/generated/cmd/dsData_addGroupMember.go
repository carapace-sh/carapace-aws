package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsData_addGroupMemberCmd = &cobra.Command{
	Use:   "add-group-member",
	Short: "Adds an existing user, group, or computer as a group member.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsData_addGroupMemberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsData_addGroupMemberCmd).Standalone()

		dsData_addGroupMemberCmd.Flags().String("client-token", "", "A unique and case-sensitive identifier that you provide to make sure the idempotency of the request, so multiple identical calls have the same effect as one single call.")
		dsData_addGroupMemberCmd.Flags().String("directory-id", "", "The identifier (ID) of the directory that's associated with the group.")
		dsData_addGroupMemberCmd.Flags().String("group-name", "", "The name of the group.")
		dsData_addGroupMemberCmd.Flags().String("member-name", "", "The `SAMAccountName` of the user, group, or computer to add as a group member.")
		dsData_addGroupMemberCmd.Flags().String("member-realm", "", "The domain name that's associated with the group member.")
		dsData_addGroupMemberCmd.MarkFlagRequired("directory-id")
		dsData_addGroupMemberCmd.MarkFlagRequired("group-name")
		dsData_addGroupMemberCmd.MarkFlagRequired("member-name")
	})
	dsDataCmd.AddCommand(dsData_addGroupMemberCmd)
}
