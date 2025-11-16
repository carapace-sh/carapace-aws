package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getJourneyExecutionActivityMetricsCmd = &cobra.Command{
	Use:   "get-journey-execution-activity-metrics",
	Short: "Retrieves (queries) pre-aggregated data for a standard execution metric that applies to a journey activity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getJourneyExecutionActivityMetricsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getJourneyExecutionActivityMetricsCmd).Standalone()

		pinpoint_getJourneyExecutionActivityMetricsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getJourneyExecutionActivityMetricsCmd.Flags().String("journey-activity-id", "", "The unique identifier for the journey activity.")
		pinpoint_getJourneyExecutionActivityMetricsCmd.Flags().String("journey-id", "", "The unique identifier for the journey.")
		pinpoint_getJourneyExecutionActivityMetricsCmd.Flags().String("next-token", "", "The `string that specifies which page of results to return in a paginated response.")
		pinpoint_getJourneyExecutionActivityMetricsCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
		pinpoint_getJourneyExecutionActivityMetricsCmd.MarkFlagRequired("application-id")
		pinpoint_getJourneyExecutionActivityMetricsCmd.MarkFlagRequired("journey-activity-id")
		pinpoint_getJourneyExecutionActivityMetricsCmd.MarkFlagRequired("journey-id")
	})
	pinpointCmd.AddCommand(pinpoint_getJourneyExecutionActivityMetricsCmd)
}
