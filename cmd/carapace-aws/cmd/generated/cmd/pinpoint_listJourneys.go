package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_listJourneysCmd = &cobra.Command{
	Use:   "list-journeys",
	Short: "Retrieves information about the status, configuration, and other settings for all the journeys that are associated with an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_listJourneysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_listJourneysCmd).Standalone()

		pinpoint_listJourneysCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_listJourneysCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
		pinpoint_listJourneysCmd.Flags().String("token", "", "The NextToken string that specifies which page of results to return in a paginated response.")
		pinpoint_listJourneysCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_listJourneysCmd)
}
