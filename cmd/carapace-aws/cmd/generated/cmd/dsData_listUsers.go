package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsData_listUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "Returns user information for the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsData_listUsersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsData_listUsersCmd).Standalone()

		dsData_listUsersCmd.Flags().String("directory-id", "", "The identifier (ID) of the directory that's associated with the user.")
		dsData_listUsersCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		dsData_listUsersCmd.Flags().String("next-token", "", "An encoded paging token for paginated calls that can be passed back to retrieve the next page.")
		dsData_listUsersCmd.Flags().String("realm", "", "The domain name that's associated with the user.")
		dsData_listUsersCmd.MarkFlagRequired("directory-id")
	})
	dsDataCmd.AddCommand(dsData_listUsersCmd)
}
