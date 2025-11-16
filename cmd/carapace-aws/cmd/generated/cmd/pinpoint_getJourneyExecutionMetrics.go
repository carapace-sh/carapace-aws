package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getJourneyExecutionMetricsCmd = &cobra.Command{
	Use:   "get-journey-execution-metrics",
	Short: "Retrieves (queries) pre-aggregated data for a standard execution metric that applies to a journey.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getJourneyExecutionMetricsCmd).Standalone()

	pinpoint_getJourneyExecutionMetricsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getJourneyExecutionMetricsCmd.Flags().String("journey-id", "", "The unique identifier for the journey.")
	pinpoint_getJourneyExecutionMetricsCmd.Flags().String("next-token", "", "The `string that specifies which page of results to return in a paginated response.")
	pinpoint_getJourneyExecutionMetricsCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
	pinpoint_getJourneyExecutionMetricsCmd.MarkFlagRequired("application-id")
	pinpoint_getJourneyExecutionMetricsCmd.MarkFlagRequired("journey-id")
	pinpointCmd.AddCommand(pinpoint_getJourneyExecutionMetricsCmd)
}
