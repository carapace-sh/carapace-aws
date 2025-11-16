package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getCampaignDateRangeKpiCmd = &cobra.Command{
	Use:   "get-campaign-date-range-kpi",
	Short: "Retrieves (queries) pre-aggregated data for a standard metric that applies to a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getCampaignDateRangeKpiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getCampaignDateRangeKpiCmd).Standalone()

		pinpoint_getCampaignDateRangeKpiCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getCampaignDateRangeKpiCmd.Flags().String("campaign-id", "", "The unique identifier for the campaign.")
		pinpoint_getCampaignDateRangeKpiCmd.Flags().String("end-time", "", "The last date and time to retrieve data for, as part of an inclusive date range that filters the query results.")
		pinpoint_getCampaignDateRangeKpiCmd.Flags().String("kpi-name", "", "The name of the metric, also referred to as a *key performance indicator (KPI)*, to retrieve data for.")
		pinpoint_getCampaignDateRangeKpiCmd.Flags().String("next-token", "", "The string that specifies which page of results to return in a paginated response.")
		pinpoint_getCampaignDateRangeKpiCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
		pinpoint_getCampaignDateRangeKpiCmd.Flags().String("start-time", "", "The first date and time to retrieve data for, as part of an inclusive date range that filters the query results.")
		pinpoint_getCampaignDateRangeKpiCmd.MarkFlagRequired("application-id")
		pinpoint_getCampaignDateRangeKpiCmd.MarkFlagRequired("campaign-id")
		pinpoint_getCampaignDateRangeKpiCmd.MarkFlagRequired("kpi-name")
	})
	pinpointCmd.AddCommand(pinpoint_getCampaignDateRangeKpiCmd)
}
