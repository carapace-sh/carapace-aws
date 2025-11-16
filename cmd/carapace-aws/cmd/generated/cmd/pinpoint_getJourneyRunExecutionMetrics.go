package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getJourneyRunExecutionMetricsCmd = &cobra.Command{
	Use:   "get-journey-run-execution-metrics",
	Short: "Retrieves (queries) pre-aggregated data for a standard run execution metric that applies to a journey.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getJourneyRunExecutionMetricsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getJourneyRunExecutionMetricsCmd).Standalone()

		pinpoint_getJourneyRunExecutionMetricsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getJourneyRunExecutionMetricsCmd.Flags().String("journey-id", "", "The unique identifier for the journey.")
		pinpoint_getJourneyRunExecutionMetricsCmd.Flags().String("next-token", "", "The `string that specifies which page of results to return in a paginated response.")
		pinpoint_getJourneyRunExecutionMetricsCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
		pinpoint_getJourneyRunExecutionMetricsCmd.Flags().String("run-id", "", "The unique identifier for the journey run.")
		pinpoint_getJourneyRunExecutionMetricsCmd.MarkFlagRequired("application-id")
		pinpoint_getJourneyRunExecutionMetricsCmd.MarkFlagRequired("journey-id")
		pinpoint_getJourneyRunExecutionMetricsCmd.MarkFlagRequired("run-id")
	})
	pinpointCmd.AddCommand(pinpoint_getJourneyRunExecutionMetricsCmd)
}
