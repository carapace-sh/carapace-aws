package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getJourneyDateRangeKpiCmd = &cobra.Command{
	Use:   "get-journey-date-range-kpi",
	Short: "Retrieves (queries) pre-aggregated data for a standard engagement metric that applies to a journey.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getJourneyDateRangeKpiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getJourneyDateRangeKpiCmd).Standalone()

		pinpoint_getJourneyDateRangeKpiCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getJourneyDateRangeKpiCmd.Flags().String("end-time", "", "The last date and time to retrieve data for, as part of an inclusive date range that filters the query results.")
		pinpoint_getJourneyDateRangeKpiCmd.Flags().String("journey-id", "", "The unique identifier for the journey.")
		pinpoint_getJourneyDateRangeKpiCmd.Flags().String("kpi-name", "", "The name of the metric, also referred to as a *key performance indicator (KPI)*, to retrieve data for.")
		pinpoint_getJourneyDateRangeKpiCmd.Flags().String("next-token", "", "The string that specifies which page of results to return in a paginated response.")
		pinpoint_getJourneyDateRangeKpiCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
		pinpoint_getJourneyDateRangeKpiCmd.Flags().String("start-time", "", "The first date and time to retrieve data for, as part of an inclusive date range that filters the query results.")
		pinpoint_getJourneyDateRangeKpiCmd.MarkFlagRequired("application-id")
		pinpoint_getJourneyDateRangeKpiCmd.MarkFlagRequired("journey-id")
		pinpoint_getJourneyDateRangeKpiCmd.MarkFlagRequired("kpi-name")
	})
	pinpointCmd.AddCommand(pinpoint_getJourneyDateRangeKpiCmd)
}
