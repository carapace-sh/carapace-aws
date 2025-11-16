package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsData_listGroupsCmd = &cobra.Command{
	Use:   "list-groups",
	Short: "Returns group information for the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsData_listGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsData_listGroupsCmd).Standalone()

		dsData_listGroupsCmd.Flags().String("directory-id", "", "The identifier (ID) of the directory that's associated with the group.")
		dsData_listGroupsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		dsData_listGroupsCmd.Flags().String("next-token", "", "An encoded paging token for paginated calls that can be passed back to retrieve the next page.")
		dsData_listGroupsCmd.Flags().String("realm", "", "The domain name associated with the directory.")
		dsData_listGroupsCmd.MarkFlagRequired("directory-id")
	})
	dsDataCmd.AddCommand(dsData_listGroupsCmd)
}
