package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchRoutingProfilesCmd = &cobra.Command{
	Use:   "search-routing-profiles",
	Short: "Searches routing profiles in an Amazon Connect instance, with optional filtering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchRoutingProfilesCmd).Standalone()

	connect_searchRoutingProfilesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_searchRoutingProfilesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_searchRoutingProfilesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_searchRoutingProfilesCmd.Flags().String("search-criteria", "", "The search criteria to be used to return routing profiles.")
	connect_searchRoutingProfilesCmd.Flags().String("search-filter", "", "Filters to be applied to search results.")
	connect_searchRoutingProfilesCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_searchRoutingProfilesCmd)
}
