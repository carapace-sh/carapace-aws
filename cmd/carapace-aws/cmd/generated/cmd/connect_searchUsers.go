package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchUsersCmd = &cobra.Command{
	Use:   "search-users",
	Short: "Searches users in an Amazon Connect instance, with optional filtering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchUsersCmd).Standalone()

	connect_searchUsersCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_searchUsersCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_searchUsersCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_searchUsersCmd.Flags().String("search-criteria", "", "")
	connect_searchUsersCmd.Flags().String("search-filter", "", "Filters to be applied to search results.")
	connect_searchUsersCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_searchUsersCmd)
}
