package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsData_listGroupMembersCmd = &cobra.Command{
	Use:   "list-group-members",
	Short: "Returns member information for the specified group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsData_listGroupMembersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsData_listGroupMembersCmd).Standalone()

		dsData_listGroupMembersCmd.Flags().String("directory-id", "", "The identifier (ID) of the directory that's associated with the group.")
		dsData_listGroupMembersCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		dsData_listGroupMembersCmd.Flags().String("member-realm", "", "The domain name that's associated with the group member.")
		dsData_listGroupMembersCmd.Flags().String("next-token", "", "An encoded paging token for paginated calls that can be passed back to retrieve the next page.")
		dsData_listGroupMembersCmd.Flags().String("realm", "", "The domain name that's associated with the group.")
		dsData_listGroupMembersCmd.Flags().String("samaccount-name", "", "The name of the group.")
		dsData_listGroupMembersCmd.MarkFlagRequired("directory-id")
		dsData_listGroupMembersCmd.MarkFlagRequired("samaccount-name")
	})
	dsDataCmd.AddCommand(dsData_listGroupMembersCmd)
}
