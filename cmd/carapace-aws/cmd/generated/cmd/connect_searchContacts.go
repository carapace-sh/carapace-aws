package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchContactsCmd = &cobra.Command{
	Use:   "search-contacts",
	Short: "Searches contacts in an Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchContactsCmd).Standalone()

	connect_searchContactsCmd.Flags().String("instance-id", "", "The identifier of Amazon Connect instance.")
	connect_searchContactsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_searchContactsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_searchContactsCmd.Flags().String("search-criteria", "", "The search criteria to be used to return contacts.")
	connect_searchContactsCmd.Flags().String("sort", "", "Specifies a field to sort by and a sort order.")
	connect_searchContactsCmd.Flags().String("time-range", "", "Time range that you want to search results.")
	connect_searchContactsCmd.MarkFlagRequired("instance-id")
	connect_searchContactsCmd.MarkFlagRequired("time-range")
	connectCmd.AddCommand(connect_searchContactsCmd)
}
