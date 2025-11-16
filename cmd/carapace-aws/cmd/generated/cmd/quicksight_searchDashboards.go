package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_searchDashboardsCmd = &cobra.Command{
	Use:   "search-dashboards",
	Short: "Searches for dashboards that belong to a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_searchDashboardsCmd).Standalone()

	quicksight_searchDashboardsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the user whose dashboards you're searching for.")
	quicksight_searchDashboardsCmd.Flags().String("filters", "", "The filters to apply to the search.")
	quicksight_searchDashboardsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	quicksight_searchDashboardsCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
	quicksight_searchDashboardsCmd.MarkFlagRequired("aws-account-id")
	quicksight_searchDashboardsCmd.MarkFlagRequired("filters")
	quicksightCmd.AddCommand(quicksight_searchDashboardsCmd)
}
