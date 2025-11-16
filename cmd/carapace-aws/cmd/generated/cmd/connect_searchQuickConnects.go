package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchQuickConnectsCmd = &cobra.Command{
	Use:   "search-quick-connects",
	Short: "Searches quick connects in an Amazon Connect instance, with optional filtering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchQuickConnectsCmd).Standalone()

	connect_searchQuickConnectsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_searchQuickConnectsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_searchQuickConnectsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_searchQuickConnectsCmd.Flags().String("search-criteria", "", "The search criteria to be used to return quick connects.")
	connect_searchQuickConnectsCmd.Flags().String("search-filter", "", "Filters to be applied to search results.")
	connect_searchQuickConnectsCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_searchQuickConnectsCmd)
}
