package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getJourneyRunExecutionActivityMetricsCmd = &cobra.Command{
	Use:   "get-journey-run-execution-activity-metrics",
	Short: "Retrieves (queries) pre-aggregated data for a standard run execution metric that applies to a journey activity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getJourneyRunExecutionActivityMetricsCmd).Standalone()

	pinpoint_getJourneyRunExecutionActivityMetricsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getJourneyRunExecutionActivityMetricsCmd.Flags().String("journey-activity-id", "", "The unique identifier for the journey activity.")
	pinpoint_getJourneyRunExecutionActivityMetricsCmd.Flags().String("journey-id", "", "The unique identifier for the journey.")
	pinpoint_getJourneyRunExecutionActivityMetricsCmd.Flags().String("next-token", "", "The `string that specifies which page of results to return in a paginated response.")
	pinpoint_getJourneyRunExecutionActivityMetricsCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
	pinpoint_getJourneyRunExecutionActivityMetricsCmd.Flags().String("run-id", "", "The unique identifier for the journey run.")
	pinpoint_getJourneyRunExecutionActivityMetricsCmd.MarkFlagRequired("application-id")
	pinpoint_getJourneyRunExecutionActivityMetricsCmd.MarkFlagRequired("journey-activity-id")
	pinpoint_getJourneyRunExecutionActivityMetricsCmd.MarkFlagRequired("journey-id")
	pinpoint_getJourneyRunExecutionActivityMetricsCmd.MarkFlagRequired("run-id")
	pinpointCmd.AddCommand(pinpoint_getJourneyRunExecutionActivityMetricsCmd)
}
