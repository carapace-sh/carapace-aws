package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_listEventTrackersCmd = &cobra.Command{
	Use:   "list-event-trackers",
	Short: "Returns the list of event trackers associated with the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_listEventTrackersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_listEventTrackersCmd).Standalone()

		personalize_listEventTrackersCmd.Flags().String("dataset-group-arn", "", "The ARN of a dataset group used to filter the response.")
		personalize_listEventTrackersCmd.Flags().String("max-results", "", "The maximum number of event trackers to return.")
		personalize_listEventTrackersCmd.Flags().String("next-token", "", "A token returned from the previous call to `ListEventTrackers` for getting the next set of event trackers (if they exist).")
	})
	personalizeCmd.AddCommand(personalize_listEventTrackersCmd)
}
