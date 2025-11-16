package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_listEventPredictionsCmd = &cobra.Command{
	Use:   "list-event-predictions",
	Short: "Gets a list of past predictions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_listEventPredictionsCmd).Standalone()

	frauddetector_listEventPredictionsCmd.Flags().String("detector-id", "", "The detector ID.")
	frauddetector_listEventPredictionsCmd.Flags().String("detector-version-id", "", "The detector version ID.")
	frauddetector_listEventPredictionsCmd.Flags().String("event-id", "", "The event ID.")
	frauddetector_listEventPredictionsCmd.Flags().String("event-type", "", "The event type associated with the detector.")
	frauddetector_listEventPredictionsCmd.Flags().String("max-results", "", "The maximum number of predictions to return for the request.")
	frauddetector_listEventPredictionsCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	frauddetector_listEventPredictionsCmd.Flags().String("prediction-time-range", "", "The time period for when the predictions were generated.")
	frauddetectorCmd.AddCommand(frauddetector_listEventPredictionsCmd)
}
