package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getJourneyRunsCmd = &cobra.Command{
	Use:   "get-journey-runs",
	Short: "Provides information about the runs of a journey.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getJourneyRunsCmd).Standalone()

	pinpoint_getJourneyRunsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getJourneyRunsCmd.Flags().String("journey-id", "", "The unique identifier for the journey.")
	pinpoint_getJourneyRunsCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
	pinpoint_getJourneyRunsCmd.Flags().String("token", "", "The NextToken string that specifies which page of results to return in a paginated response.")
	pinpoint_getJourneyRunsCmd.MarkFlagRequired("application-id")
	pinpoint_getJourneyRunsCmd.MarkFlagRequired("journey-id")
	pinpointCmd.AddCommand(pinpoint_getJourneyRunsCmd)
}
