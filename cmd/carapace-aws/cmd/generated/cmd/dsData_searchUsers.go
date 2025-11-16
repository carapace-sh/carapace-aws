package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsData_searchUsersCmd = &cobra.Command{
	Use:   "search-users",
	Short: "Searches the specified directory for a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsData_searchUsersCmd).Standalone()

	dsData_searchUsersCmd.Flags().String("directory-id", "", "The identifier (ID) of the directory that's associated with the user.")
	dsData_searchUsersCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	dsData_searchUsersCmd.Flags().String("next-token", "", "An encoded paging token for paginated calls that can be passed back to retrieve the next page.")
	dsData_searchUsersCmd.Flags().String("realm", "", "The domain name that's associated with the user.")
	dsData_searchUsersCmd.Flags().String("search-attributes", "", "One or more data attributes that are used to search for a user.")
	dsData_searchUsersCmd.Flags().String("search-string", "", "The attribute value that you want to search for.")
	dsData_searchUsersCmd.MarkFlagRequired("directory-id")
	dsData_searchUsersCmd.MarkFlagRequired("search-attributes")
	dsData_searchUsersCmd.MarkFlagRequired("search-string")
	dsDataCmd.AddCommand(dsData_searchUsersCmd)
}
