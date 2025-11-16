package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsData_listGroupsForMemberCmd = &cobra.Command{
	Use:   "list-groups-for-member",
	Short: "Returns group information for the specified member.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsData_listGroupsForMemberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsData_listGroupsForMemberCmd).Standalone()

		dsData_listGroupsForMemberCmd.Flags().String("directory-id", "", "The identifier (ID) of the directory that's associated with the member.")
		dsData_listGroupsForMemberCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		dsData_listGroupsForMemberCmd.Flags().String("member-realm", "", "The domain name that's associated with the group member.")
		dsData_listGroupsForMemberCmd.Flags().String("next-token", "", "An encoded paging token for paginated calls that can be passed back to retrieve the next page.")
		dsData_listGroupsForMemberCmd.Flags().String("realm", "", "The domain name that's associated with the group.")
		dsData_listGroupsForMemberCmd.Flags().String("samaccount-name", "", "The `SAMAccountName` of the user, group, or computer that's a member of the group.")
		dsData_listGroupsForMemberCmd.MarkFlagRequired("directory-id")
		dsData_listGroupsForMemberCmd.MarkFlagRequired("samaccount-name")
	})
	dsDataCmd.AddCommand(dsData_listGroupsForMemberCmd)
}
