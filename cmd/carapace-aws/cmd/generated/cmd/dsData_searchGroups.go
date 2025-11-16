package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsData_searchGroupsCmd = &cobra.Command{
	Use:   "search-groups",
	Short: "Searches the specified directory for a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsData_searchGroupsCmd).Standalone()

	dsData_searchGroupsCmd.Flags().String("directory-id", "", "The identifier (ID) of the directory that's associated with the group.")
	dsData_searchGroupsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	dsData_searchGroupsCmd.Flags().String("next-token", "", "An encoded paging token for paginated calls that can be passed back to retrieve the next page.")
	dsData_searchGroupsCmd.Flags().String("realm", "", "The domain name that's associated with the group.")
	dsData_searchGroupsCmd.Flags().String("search-attributes", "", "One or more data attributes that are used to search for a group.")
	dsData_searchGroupsCmd.Flags().String("search-string", "", "The attribute value that you want to search for.")
	dsData_searchGroupsCmd.MarkFlagRequired("directory-id")
	dsData_searchGroupsCmd.MarkFlagRequired("search-attributes")
	dsData_searchGroupsCmd.MarkFlagRequired("search-string")
	dsDataCmd.AddCommand(dsData_searchGroupsCmd)
}
