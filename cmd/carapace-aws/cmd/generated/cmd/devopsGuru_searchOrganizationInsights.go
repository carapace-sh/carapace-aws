package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_searchOrganizationInsightsCmd = &cobra.Command{
	Use:   "search-organization-insights",
	Short: "Returns a list of insights in your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_searchOrganizationInsightsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_searchOrganizationInsightsCmd).Standalone()

		devopsGuru_searchOrganizationInsightsCmd.Flags().String("account-ids", "", "The ID of the Amazon Web Services account.")
		devopsGuru_searchOrganizationInsightsCmd.Flags().String("filters", "", "A `SearchOrganizationInsightsFilters` object that is used to set the severity and status filters on your insight search.")
		devopsGuru_searchOrganizationInsightsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		devopsGuru_searchOrganizationInsightsCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
		devopsGuru_searchOrganizationInsightsCmd.Flags().String("start-time-range", "", "")
		devopsGuru_searchOrganizationInsightsCmd.Flags().String("type", "", "The type of insights you are searching for (`REACTIVE` or `PROACTIVE`).")
		devopsGuru_searchOrganizationInsightsCmd.MarkFlagRequired("account-ids")
		devopsGuru_searchOrganizationInsightsCmd.MarkFlagRequired("start-time-range")
		devopsGuru_searchOrganizationInsightsCmd.MarkFlagRequired("type")
	})
	devopsGuruCmd.AddCommand(devopsGuru_searchOrganizationInsightsCmd)
}
