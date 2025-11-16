package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_searchInsightsCmd = &cobra.Command{
	Use:   "search-insights",
	Short: "Returns a list of insights in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_searchInsightsCmd).Standalone()

	devopsGuru_searchInsightsCmd.Flags().String("filters", "", "A `SearchInsightsFilters` object that is used to set the severity and status filters on your insight search.")
	devopsGuru_searchInsightsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	devopsGuru_searchInsightsCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	devopsGuru_searchInsightsCmd.Flags().String("start-time-range", "", "The start of the time range passed in.")
	devopsGuru_searchInsightsCmd.Flags().String("type", "", "The type of insights you are searching for (`REACTIVE` or `PROACTIVE`).")
	devopsGuru_searchInsightsCmd.MarkFlagRequired("start-time-range")
	devopsGuru_searchInsightsCmd.MarkFlagRequired("type")
	devopsGuruCmd.AddCommand(devopsGuru_searchInsightsCmd)
}
